# Go Routine 정리

- Golang 은 뛰어난 동시성 지원을 장점으로 가지고 있습니다.
- Golang 에서는 Thread 가 아닌 GoRoutine 을 활용하여 동시성 프로그래밍 개발을 진행합니다.

## 고루틴 (goroutine)

- “lightweight thread managed by the Go runtime”
- Go 런타임에서 관리되는 경량 쓰레드
- OS 레벨의 쓰레드가 아닌 golang 런타임에서 관리되는 논리적 쓰레드

```go
package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)

	counter := 0

	// goroutine 생성 및 다음 익명 함수의 작업을 할당
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			counter++
		}
	}()

	wg.Wait() // goroutine 작업의 완료까지 대기
	fmt.Println("counter :", counter)

}
```

## goroutine vs Thread

- 메모리 소비
  - golang 에서 고루틴을 생성할 때는 많은 메모리를 필요로 하지 않습니다.
  - 2KB 의 스택 메모리 공간만 필요로 하며, 필요에 따라 힙 메모리 공간을 사용하기도 합니다.
  - 반면에 쓰레드는 Guard Page 라고 불리는 메모리 영역과 함께 포함하여 약 1MB 의 메모리 공간을 소모하여 생성됩니다.
- 생성, 소거 비용
  - OS 쓰레드는 생성/소거시 많은 비용을 소모합니다.
  - golang 의 고루틴은 런타임에서 하드웨어에 의존하는 것이 아니라 논리적으로 생성/소거되기 때문에 상대적으로 소모되는 비용이 저렴합니다.
  - 고루틴은 OS 쓰레드를 사용하는 언어들과 비교하여 상대적으로 OOM 이슈의 발생 가능성이 적습니다.
- Context Switching 비용
  - 하나의 쓰레드가 작업을 처리하기 위해 Blocking 된다면 다른 쓰레드가 새로운 작업을 처리하도록 스케쥴링되어 있습니다.
  - 쓰레드가 스케쥴링되고 교체되는 동안에 스케쥴러에서 모든 레지스터들은 save/restore 해야 합니다.
  - 이 경우 golang 의 고루틴은 3개의 레지스터 (Program Counter, Stack Pointer, DX)만 save/restore 작업을 하기 때문에 OS 쓰레드를 사용하는 경우보다 상대적으로 Context Switching 비용을 적게 소모합니다.

## 고루틴은 어떻게 실행될까?

- golang 은 프로그램 시작부터 끝나는 시점까지 런타임 내내 고루틴을 관리합니다.
- 고루틴은 M:N 쓰레드 모델 (LWP) 을 채택하고 있어 기존의 Thread/Thread Pool 을 활용하는 방식보다 더 가볍고 빠른 특성을 지니고 있습니다.

## 쓰레드의 종류

- User-Level 쓰레드 -사용자 라이브러리를 통해 쓰레드 관리 기능이 제공되며, 여러 쓰레드가 1개의 OS 쓰레드 위에서 동작하는 형태 (1:N)
  - Context Switching 비용이 적기 때문에 속도가 빠릅니다.
  - OS 쓰레드 1개만 사용하는 구조이기 때문에 멀티 코어 프로그래밍이 불가능합니다.
- Kernel-Level 쓰레드
  - 운영체제가 지원하는 기능으로 구현되어 Kernel 이 쓰레드의 생성, 스케쥴링을 담당합니다. (1:1)
  - Kernel 은 프로세스 내의 다른 쓰레드를 중단시키지 않고 계속 실행됩니다.
  - 멀티 프로세싱 환경에서는 커널은 여러개의 쓰레드르 각각 다른 프로세서에 할당할 수 있습니다.
  - Context Switching 비용이 크기 때문에 속도가 느립니다.
- Combined (golang goroutine 채택 방식)
  - Kernel 쓰레드와 User 쓰레드를 혼합하여 사용하는 방식입니다. (M:N)
  - User-Level 쓰레드는 LWP 에 의해 다중화 됩니다.
  - LWP 는 Kernel 과 프로세스 사이에서 중간자 역할을 합니다.
  - Context Switching 속도가 빠르고, 멀티 코어 활용도 가능합니다.

## Go 스케쥴러

![GMP 모델](./GMP%20%EB%AA%A8%EB%8D%B8.png)

- golang 의 스케쥴러는 G-M-P 모델로 표현되어 스케쥴링이 처리되는 구조입니다.
- G (Goroutine)
  - 고루틴을 의미하며, 고루틴을 구성하는 논리적 구조체의 구현체를 의미합니다.
  - Context Switching 을 위해 SP, goroutine 의 상태 정보 등을 가지고 있습니다.
  - G 는 LRQ 에서 대기하고 있습니다.
- M (Machine)
  - OS 쓰레드를 의미하며, 실제 OS 쓰레드가 아닌 논리적 구현체로 표준 POSIX 쓰레드를 따릅니다.
  - M 은 P 로부터 G 를 할당받아 실행합니다.
  - goroutine 과 OS 쓰레드를 연결하므로 쓰레드 핸들 정보, 실행중인 고루틴, P 의 포인터를 가지고 있습니다.
- P (Processor)
  - 프로세스를 의미하며, 논리적인 프로세스로 스케쥴링과 관련된 Context 정보를 가지고 있습니다.
  - 런타임시 GOMAXPROCS 설정 값만큼의 개수로 프로세서를 가질 수 있습니다.
  - P 는 Context 정보를 담고 있으며, P:LRQ = 1:1 비율이며, G 를 M 에 할당하는 역할을 수행합니다.
- LRQ (LocalRunQueue)
  - P 에 종속되어 있는 Run Queue 이며, LRQ 에 실행 가능한 고루틴들이 적재됩니다.
  - P 는 LRQ 로부터 고루틴을 하나씩 pop 하고, M 에 할당하여 고루틴을 실행합니다.
- GRQ (GlobalRunQueue)
  - LRQ 에 할당되지 못한 고루틴을 관리하는 Run Queue
  - 실행 상태의 고루틴은 한번에 최대 10ms 까지 실행되며, 실행된 고루틴은 대기 상태로 변하고 GRQ 로 적재됩니다.
  - 고루틴이 생성되는 시점에 모든 LRQ 가 가득찬 경우에도 GRQ 에 고루틴이 적재됩니다.

## 스케쥴러 작동 원리

![GMP 스케쥴러 작동 원리](./GMP%20%EC%8A%A4%EC%BC%80%EC%A5%B4%EB%9F%AC.png)

- syscall
  - 로직을 실행하던 도중에 syscall(주로 I/O 작업 같은 행위 등)이 발생하게 되면 blocking이 발생되는데, 이 현상이 해당 작업을 처리하던 쓰레드에 영향을 주어 다음 작업을 처리할 수 없기 때문에 성능 저하의 원인이 될 수 있다. Go에서는 스케줄러가 작업을 멈추지 않고 계속 진행할 수 있도록 syscall이 발생한 고루틴을 다른 쓰레드로 넘기고 P의 LRQ에 적재되어 있던 다음 고루틴이 정상적으로 처리될 수 있도록 보장한다. 이후 syscall 처리가 끝난 고루틴은 잠시 넘겨주었던 P로 다시 적재되거나 GRQ에 적재된다.
  - 다수의 고루틴은 소수의 쓰레드 위에서 동작한다. 즉 다중화(Multiplexing)되어 돌아간다고 앞서 언급 했었는데, 바로 위에서 설명한 과정의 Go 스케줄러의 스케줄링 덕분에 가능해진 것이다. 따라서 M에 바로 P(Context)가 붙어서 고루틴 간에 Context Switching이 발생하게 되는 구조로 여기서 P(Context)의 개수는 Go의 GOMAXPROCS라는 환경변수 값을 통해 조정이 가능하다.
- Work Stealing
  - M, P(Context)가 모든 작업을 마치게 되면 먼저 GRQ에 쌓여있는 고루틴을 가져오려 시도하고 이마저도 없으면 다른 P의 LRQ에서 절반의 고루틴(작업)을 가져온다(Steal). 작업의 불균형으로 인한 병목현상을 이러한 Work Stealing 기법을 통해 리소스를 더 효율적으로 사용하게 된다.

## 고루틴의 Context Switching 시점

- unbuffered channel에 접근할 때(write or read)
- 시스템 I/O가 발생했을 때
- 메모리가 할당될 때
- time.Sleep() Function 같은 해당 문맥에서 sleep 처리하는 로직이 실행될 때
- runtime.Gosched() Function 같은 로직이 실행될 때

## 용어

- OOM (OutOfMemory)

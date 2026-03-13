package main

/*
Rules:

1. Philosophers are numbered 0 to 4.
2. Philosopher i picks up the fork on the left (i) and the fork on the right ((i+1) % 5).
3. No two philosophers should hold the same fork simultaneously.
4. The execution order of WantsToEat calls may be arbitrary (they may run in parallel).
5. You must ensure no deadlocks happen.
*/

func main() {
	
}

type DiningPhilosophers struct {
    // your fields here
}

func Constructor() DiningPhilosophers {
    return DiningPhilosophers{}
}

func (dp *DiningPhilosophers) WantsToEat(
    philosopher int,
    pickLeftFork func(),
    pickRightFork func(),
    eat func(),
    putLeftFork func(),
    putRightFork func(),
) {
    // implement logic
}

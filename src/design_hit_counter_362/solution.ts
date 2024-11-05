class HitCounter {
  private hits: number[]

  constructor() {
    this.hits = []
  }

  hit(timestamp: number): void {
    this.hits.push(timestamp)
  }

  getHits(timestamp: number): number {
    this.cleanOldHits(timestamp)
    return this.hits.length
  }

  private cleanOldHits(currentTime: number): void {
    const threshold = currentTime - 300
    while (this.hits.length > 0 && this.hits[0] <= threshold) {
      this.hits.shift()
    }
  }
}

export function solution(): void {
  const counter = new HitCounter()
  counter.hit(1)
  counter.hit(2)
  counter.hit(3)
  console.log(counter.getHits(4))
  counter.hit(300)
  console.log(counter.getHits(300))
  console.log(counter.getHits(301))
}

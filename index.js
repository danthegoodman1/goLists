const BPlusTree = require('bplustree')

const main = async () => {
  const tree = new BPlusTree()
  const start = new Date()
  for (let i = 0; i < 100000000; i++) {
    tree.store(i, i)
  }
  const end = new Date()
  console.log(start.getTime())
  console.log(end.getTime())
  console.log("CONSTRUCTION TIME:", end.getTime() - start.getTime())

  const s2 = process.hrtime.bigint()
  tree.fetch(453012)
  const e2 = process.hrtime.bigint()
  console.log("FETCH TIME:", e2 - s2)
}

main()

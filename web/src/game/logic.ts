export type Board = number[][]
export type Dir = 'left'|'right'|'up'|'down'

export function makeBoard(): Board { return Array.from({length:4},()=>[0,0,0,0]) }

export function clone(b: Board): Board { return b.map(r=>r.slice()) }

export function addRandomTile(b: Board): [number,number]|null {
  const empties: [number,number][] = []
  for (let i=0;i<4;i++) for (let j=0;j<4;j++) if (b[i][j]===0) empties.push([i,j])
  if (!empties.length) return null
  const [i,j] = empties[Math.floor(Math.random()*empties.length)]
  b[i][j] = Math.random()<0.9?2:4
  return [i,j]
}

export function startGame(): { board: Board, score: number, newTiles: [number,number][] }{
  const board = makeBoard()
  const newTiles: [number,number][] = []
  const t1 = addRandomTile(board); if(t1) newTiles.push(t1)
  const t2 = addRandomTile(board); if(t2) newTiles.push(t2)
  return { board, score: 0, newTiles }
}

function compress(row: number[]): { row:number[], gained:number }{
  const arr = row.filter(v=>v!==0)
  let gained = 0
  for (let i=0;i<arr.length-1;i++){
    if (arr[i]!==0 && arr[i]===arr[i+1]){ arr[i]*=2; gained+=arr[i]; arr[i+1]=0; i++ }
  }
  const newRow = arr.filter(v=>v!==0)
  while (newRow.length<4) newRow.push(0)
  return { row:newRow, gained }
}

function rotateR(b: Board): Board { // rotate clockwise
  const n = 4
  const r = makeBoard()
  for (let i=0;i<n;i++) for (let j=0;j<n;j++) r[j][n-1-i] = b[i][j]
  return r
}

export function move(b: Board, dir: Dir): {
  board: Board,
  moved: boolean,
  gained: number,
  merged: [number,number][],
  newTiles: [number,number][],
  moves: {from:[number,number],to:[number,number]}[]
} {
  let work = clone(b)
  let times = 0
  if (dir==='up') times = 3
  if (dir==='right') times = 2
  if (dir==='down') times = 1
  for (let k=0;k<times;k++) work = rotateR(work)

  let moved = false, total = 0, merged: [number,number][] = []
  const moves: {from:[number,number],to:[number,number]}[] = []
  // 记录每行移动轨迹
  for (let i=0;i<4;i++){
    const old = work[i].slice()
    const { row, gained } = compress(work[i])
    if (!arraysEq(row, work[i])) moved = true
    // 记录合并位置
    for(let j=0;j<4;j++) if(row[j]!==0 && row[j]!==work[i][j] && row[j]===work[i][j]*2) merged.push([i,j])
    // 记录移动轨迹
    for(let j=0;j<4;j++){
      if(old[j]!==0){
        const to = row.indexOf(old[j])
        if(to!==-1 && to!==j) moves.push({from:[i,j],to:[i,to]})
      }
    }
    work[i] = row
    total += gained
  }
  for (let k=0;k<(4-times)%4;k++) work = rotateR(work)
  // 新增方块
  const newTiles: [number,number][] = []
  if(moved){
    const t = addRandomTile(work); if(t) newTiles.push(t)
  }
  return { board: work, moved, gained: total, merged, newTiles, moves }
}

function arraysEq(a:number[], b:number[]): boolean { return a.length===b.length && a.every((v,i)=>v===b[i]) }

export function hasMoves(b: Board): boolean {
  for (let i=0;i<4;i++) for (let j=0;j<4;j++){
    if (b[i][j]===0) return true
    if (i<3 && b[i][j]===b[i+1][j]) return true
    if (j<3 && b[i][j]===b[i][j+1]) return true
  }
  return false
}

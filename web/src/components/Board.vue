<template>
  <div class="board" @touchstart="onTouchStart" @touchend="onTouchEnd">
    <div class="row" v-for="(row,i) in board" :key="i">
      <div
        v-for="(v,j) in row"
        :key="j"
        class="cell"
        :class="cellClass(i,j,v)"
        :style="tileStyle(i,j)"
      >{{ v || '' }}</div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { computed } from 'vue'
import type { Board } from '../game/logic'

const props = defineProps<{ board: Board, merged?: [number,number][], newTiles?: [number,number][], moves?: {from:[number,number],to:[number,number]}[] }>()

function cellClass(i:number, j:number, v:number){
  const c = v===0 ? 'c0' : `c${Math.min(v, 2048)}`
  let arr = [c]
  if(props.merged?.some(([x,y])=>x===i&&y===j)) arr.push('merged')
  if(props.newTiles?.some(([x,y])=>x===i&&y===j)) arr.push('new')
  if(isMoving(i,j)) arr.push('moving')
  return arr
}

function isMoving(i:number, j:number){
  return props.moves?.some(m=>m.to[0]===i&&m.to[1]===j)
}

function tileStyle(i:number, j:number){
  if(!props.moves) return {}
  const move = props.moves.find(m=>m.to[0]===i&&m.to[1]===j)
  if(move){
    const dx = (move.from[1]-move.to[1])*82
    const dy = (move.from[0]-move.to[0])*82
    return {
      transform: `translate(${dx}px,${dy}px)`,
      transition: 'transform .28s cubic-bezier(.4,1.4,.6,1)',
    }
  }
  return {}
}

let sx=0, sy=0
function onTouchStart(e:TouchEvent){ const t = e.touches[0]; sx=t.clientX; sy=t.clientY }
function onTouchEnd(e:TouchEvent){
  const t = e.changedTouches[0];
  const dx = t.clientX - sx, dy = t.clientY - sy
  const adx = Math.abs(dx), ady = Math.abs(dy)
  if (Math.max(adx,ady) < 20) return
  const ev = new CustomEvent('swipe', { detail: { dir: adx>ady ? (dx>0?'right':'left') : (dy>0?'down':'up') } })
  (e.currentTarget as HTMLElement).dispatchEvent(ev)
}
</script>
<style scoped>
/* 动画增强 */
.board{ background:#bbada0; border-radius:10px; padding:10px; user-select:none; }
.row{ display:grid; grid-template-columns: repeat(4, 1fr); gap:10px; margin-bottom:10px }
.row:last-child{ margin-bottom:0 }
.cell{ height:72px; display:flex; align-items:center; justify-content:center; font-weight:700; font-size:22px; border-radius:8px; background:#cdc1b4; color:#776e65; transition:background .2s, color .2s, transform .2s; will-change:transform; }
.c0{ background:#cdc1b4 }
.c2{ background:#eee4da }
.c4{ background:#ede0c8 }
.c8{ background:#f2b179; color:white }
.c16{ background:#f59563; color:white }
.c32{ background:#f67c5f; color:white }
.c64{ background:#f65e3b; color:white }
.c128{ background:#edcf72 }
.c256{ background:#edcc61 }
.c512{ background:#edc850 }
.c1024{ background:#edc53f }
.c2048{ background:#edc22e }
.merged{ animation:mergeFlash .35s }
.new{ animation:popIn .35s }
@keyframes mergeFlash{ 0%{box-shadow:0 0 0 0 #fbbf24; background:#fbbf24} 60%{box-shadow:0 0 16px 8px #fbbf24; background:#fbbf24} 100%{box-shadow:none} }
@keyframes popIn{ 0%{transform:scale(.2)} 80%{transform:scale(1.1)} 100%{transform:scale(1)} }
</style>

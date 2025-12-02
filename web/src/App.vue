<template>
  <div>
    <div class="topbar">
      <div>
        <h2 style="margin:0">2048</h2>
        <div class="muted">欢迎，{{ name || '游客' }}
          <button class="btn small" style="margin-left:8px" @click="showName = true">修改姓名</button>
        </div>
      </div>
      <div style="display:flex; gap:8px; align-items:center">
        <div class="card" style="padding:8px 12px">分数：<b>{{ score }}</b></div>
        <button class="btn" @click="restart">重新开始</button>
      </div>
    </div>

    <div class="tabs" style="margin-bottom:12px">
      <div class="tab" :class="{active:view==='game'}" @click="view='game'">游戏</div>
      <div class="tab" :class="{active:view==='board'}" @click="view='board'">排行/最近</div>
    </div>

    <div v-if="view==='game'" class="card">
      <p class="muted" style="margin-top:0">使用方向键或滑动来移动方块。</p>
      <Board :board="board" :merged="merged" :newTiles="newTiles" @swipe="onSwipe"/>
      <div style="display:flex; gap:8px; margin-top:12px;">
        <button class="btn secondary" @click="view='board'">查看排行</button>
      </div>
    </div>
    <div v-else>
      <Leaderboard />
    </div>

    <NameModal :show="showName" @ok="saveName" />
    <CommentModal :show="showComment" @ok="submitComment" @cancel="skipComment" />
  </div>
</template>
<script setup lang="ts">
import { onMounted, onUnmounted, ref } from 'vue'
import Board from './components/Board.vue'
import Leaderboard from './components/Leaderboard.vue'
import NameModal from './components/NameModal.vue'
import CommentModal from './components/CommentModal.vue'
import { startGame, move, hasMoves, type Board as B } from './game/logic'
import { submitScore } from './api'

const view = ref<'game'|'board'>('game')
const name = ref('')
const score = ref(0)
const board = ref<B>([[0,0,0,0],[0,0,0,0],[0,0,0,0],[0,0,0,0]])
const showName = ref(false)
const showComment = ref(false)
const merged = ref<[number,number][]>([])
const newTiles = ref<[number,number][]>([])

function newGame(){
  const g = startGame(); board.value = g.board; score.value = g.score; merged.value=[]; newTiles.value=g.newTiles
}

function restart(){ newGame() }

function onKey(e:KeyboardEvent){
  const map: Record<string,'left'|'right'|'up'|'down'> = { ArrowLeft:'left', ArrowRight:'right', ArrowUp:'up', ArrowDown:'down' }
  const dir = map[e.key]; if(!dir) return
  doMove(dir)
}

function onSwipe(ev:CustomEvent){ doMove((ev as any).detail.dir) }

function doMove(dir:'left'|'right'|'up'|'down'){
  const { board: nb, moved, gained, merged: m, newTiles: nt } = move(board.value, dir)
  if(!moved) return
  board.value = nb
  score.value += gained
  merged.value = m
  newTiles.value = nt
  setTimeout(()=>{ merged.value=[]; newTiles.value=[] }, 350)
  if(!hasMoves(board.value)){
    showComment.value = true
  }
}

function saveName(n:string){
  name.value = n
  localStorage.setItem('nickname', n)
  showName.value = false
  newGame()
}

async function submitComment(c:string){
  showComment.value = false
  try{
    await submitScore({ name: name.value, score: score.value, comment: c })
    view.value = 'board'
  }catch(err){ alert('提交失败，请稍后重试') }
}

function skipComment(){ showComment.value = false }

onMounted(()=>{
  const n = localStorage.getItem('nickname')
  if(!n){ showName.value = true } else { name.value=n; newGame() }
  window.addEventListener('keydown', onKey)
})

onUnmounted(()=>{ window.removeEventListener('keydown', onKey) })
</script>

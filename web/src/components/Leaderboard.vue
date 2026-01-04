<template>
  <div class="card">
    <div class="tabs">
      <div class="tab" :class="{active:tab==='rank'}" @click="tab='rank'">分数排行</div>
      <div class="tab" :class="{active:tab==='recent'}" @click="tab='recent'">最近完成</div>
    </div>
    <div v-if="loading" class="muted" style="margin-top:8px">加载中…</div>
    <div v-else>
      <table class="table">
        <thead>
          <tr><th style="width:56px">#</th><th style="width:160px">称呼</th><th>分数</th><th style="width:200px">日期</th><th>评论</th></tr>
        </thead>
        <tbody>
          <tr v-for="(r,idx) in rows" :key="r.id">
            <td>{{ idx+1 }}</td>
            <td>{{ r.name }}</td>
            <td>{{ r.score }}</td>
            <td>{{ fmt(r.finished) }}</td>
            <td>{{ r.comment }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, watchEffect } from 'vue'
import { getRank, getRecent, type RecordVO } from '../api'

const tab = ref<'rank'|'recent'>('rank')
const rows = ref<RecordVO[]>([])
const loading = ref(false)

watchEffect(async () => {
  loading.value = true
  try{
    rows.value = tab.value==='rank' ? await getRank(50) : await getRecent(50)
  } finally { loading.value = false }
})

function fmt(s:string){
  if(!s) return s
  // 后端返回格式 `YYYY-MM-DD HH:MM:SS`，手动解析为本地时间
  const m = s.match(/^(\d{4})-(\d{2})-(\d{2}) (\d{2}):(\d{2}):(\d{2})$/)
  if(m){
    const y = Number(m[1]), mo = Number(m[2]) - 1, d = Number(m[3])
    const hh = Number(m[4]), mm = Number(m[5]), ss = Number(m[6])
    return new Date(y, mo, d, hh, mm, ss).toLocaleString()
  }
  try{ return new Date(s).toLocaleString() }catch{ return s }
}
</script>
<style scoped>
.table{ width:100%; border-collapse: collapse; margin-top:8px }
th,td{ text-align:left; border-bottom:1px solid #e5e7eb; padding:8px 6px; font-size:14px }
</style>

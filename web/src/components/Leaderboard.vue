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
  try{ return new Date(s).toLocaleString() }catch{ return s }
}
</script>
<style scoped>
.table{ width:100%; border-collapse: collapse; margin-top:8px }
th,td{ text-align:left; border-bottom:1px solid #e5e7eb; padding:8px 6px; font-size:14px }
</style>

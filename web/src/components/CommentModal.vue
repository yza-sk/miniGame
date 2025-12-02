<template>
  <div v-if="show" class="overlay">
    <div class="dialog animated">
      <h3>为本次成绩留下评论（可为空，最多40字）</h3>
      <textarea v-model.trim="localComment" rows="4" placeholder="随便说点什么吧~" maxlength="40"></textarea>
      <div class="actions">
        <button class="btn secondary" @click="cancel">跳过</button>
        <button class="btn" @click="confirm">提交</button>
      </div>
      <p v-if="localComment.length>40" class="err">评论不能超过40字</p>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, watch } from 'vue'

const props = defineProps<{ show: boolean }>()
const emits = defineEmits<{ (e:'ok', comment:string):void, (e:'cancel'):void }>()

const localComment = ref('')
watch(()=>props.show, v=>{ if(v) localComment.value='' })

function confirm(){
  if(localComment.value.length>40) return
  emits('ok', localComment.value)
}
function cancel(){ emits('cancel') }
</script>
<style scoped>
.overlay{ position:fixed; inset:0; background:rgba(15,23,42,.6); display:flex; align-items:center; justify-content:center; z-index:99; }
.dialog{ width:min(96vw, 520px); max-width:98vw; background:#fff; border-radius:12px; padding:16px; box-shadow:0 10px 30px rgba(0,0,0,.2); word-break:break-all; }
textarea{ width:100%; padding:10px 12px; border:1px solid #e5e7eb; border-radius:8px; margin:10px 0; resize:vertical; font-size:15px; box-sizing:border-box; }
.actions{ display:flex; justify-content:flex-end; gap:8px }
.btn{ background:#4f46e5; color:#fff; border:none; padding:8px 12px; border-radius:8px; cursor:pointer }
.btn.secondary{ background:#e2e8f0; color:#111827 }
.err{ color:#ef4444; font-size:12px; margin-top:4px; }
.animated{ animation:fadeIn .5s }
@keyframes fadeIn{ from{opacity:0;transform:scale(.9)} to{opacity:1;transform:scale(1)} }
</style>

<template>
  <div v-if="show" class="overlay">
    <div class="dialog animated" :class="{ shake: err }">
      <h3>请输入您的称呼</h3>
      <input v-model.trim="localName" :placeholder="placeholder" @keyup.enter="confirm" maxlength="12" />
      <div class="actions">
        <button class="btn" @click="confirm">开始游戏</button>
      </div>
      <p v-if="err" class="err">称呼不能为空且不超过12字</p>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, watch } from 'vue'

const props = defineProps<{ show: boolean, placeholder?: string }>()
const emits = defineEmits<{ (e:'ok', name:string):void }>()

const localName = ref('')
const err = ref(false)

watch(()=>props.show, v=>{ if(v){ localName.value=''; err.value=false } })

function confirm(){
  if(!localName.value || localName.value.length>12){ err.value = true; setTimeout(()=>err.value=false, 600); return }
  emits('ok', localName.value)
}
</script>
<style scoped>
.overlay{ position:fixed; inset:0; background:rgba(15,23,42,.6); display:flex; align-items:center; justify-content:center; z-index:99; }
.dialog{ width:min(96vw, 420px); max-width:98vw; background:#fff; border-radius:12px; padding:16px; box-shadow:0 10px 30px rgba(0,0,0,.2); word-break:break-all; }
input{ width:100%; padding:10px 12px; border:1px solid #e5e7eb; border-radius:8px; margin:10px 0; font-size:16px; box-sizing:border-box; }
.actions{ display:flex; justify-content:flex-end; gap:8px }
.btn{ background:#4f46e5; color:#fff; border:none; padding:8px 12px; border-radius:8px; cursor:pointer }
.err{ color:#ef4444; font-size:12px; margin-top:4px; }
.animated{ animation:fadeIn .5s }
.shake{ animation:shake .3s }
@keyframes fadeIn{ from{opacity:0;transform:scale(.9)} to{opacity:1;transform:scale(1)} }
@keyframes shake{ 0%{transform:translateX(0)} 20%{transform:translateX(-8px)} 40%{transform:translateX(8px)} 60%{transform:translateX(-8px)} 80%{transform:translateX(8px)} 100%{transform:translateX(0)} }
</style>

<template>
  <div class="debugger">
    <h3>API调试工具</h3>
    
    <div class="section">
      <h4>测试提交成绩</h4>
      <div class="form-group">
        <input v-model="testName" placeholder="测试用户名" />
        <input v-model.number="testScore" type="number" placeholder="测试分数" />
        <input v-model="testComment" placeholder="测试评论" />
        <button @click="testSubmit" class="btn">提交测试</button>
      </div>
      <div v-if="submitResult" class="result" :class="{ error: submitResult.error }">
        {{ submitResult.message }}
      </div>
    </div>

    <div class="section">
      <h4>测试获取排行榜</h4>
      <button @click="testGetRank" class="btn">获取排行榜</button>
      <div v-if="rankResult" class="result">
        <div v-if="rankResult.error" class="error">{{ rankResult.message }}</div>
        <div v-else>
          <p>获取到 {{ rankResult.data?.length }} 条记录</p>
          <ul>
            <li v-for="item in rankResult.data?.slice(0, 5)" :key="item.id">
              {{ item.name }} - {{ item.score }}分 - {{ formatDate(item.finished) }}
            </li>
          </ul>
        </div>
      </div>
    </div>

    <div class="section">
      <h4>测试获取最近记录</h4>
      <button @click="testGetRecent" class="btn">获取最近记录</button>
      <div v-if="recentResult" class="result">
        <div v-if="recentResult.error" class="error">{{ recentResult.message }}</div>
        <div v-else>
          <p>获取到 {{ recentResult.data?.length }} 条记录</p>
          <ul>
            <li v-for="item in recentResult.data?.slice(0, 5)" :key="item.id">
              {{ item.name }} - {{ item.score }}分 - {{ formatDate(item.finished) }}
            </li>
          </ul>
        </div>
      </div>
    </div>

    <div class="section">
      <h4>API基础地址</h4>
      <p>{{ apiBase }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { submitScore, getRank, getRecent, type Grade } from '../api'

const apiBase = ref((import.meta as any).env?.VITE_API_BASE || '/')

// 测试数据
const testName = ref('测试用户')
const testScore = ref(2048)
const testComment = ref('这是一个测试评论')

// 测试结果
const submitResult = ref<{ error: boolean; message: string } | null>(null)
const rankResult = ref<{ error: boolean; message: string; data?: Grade[] } | null>(null)
const recentResult = ref<{ error: boolean; message: string; data?: Grade[] } | null>(null)

// 测试提交成绩
async function testSubmit() {
  submitResult.value = null
  try {
    await submitScore({
      name: testName.value,
      score: testScore.value,
      comment: testComment.value
    })
    submitResult.value = { error: false, message: '✅ 提交成功！' }
  } catch (error) {
    submitResult.value = { 
      error: true, 
      message: `❌ 提交失败: ${error instanceof Error ? error.message : String(error)}` 
    }
  }
}

// 测试获取排行榜
async function testGetRank() {
  rankResult.value = null
  try {
    const data = await getRank(10)
    rankResult.value = { error: false, message: '', data }
  } catch (error) {
    rankResult.value = { 
      error: true, 
      message: `❌ 获取失败: ${error instanceof Error ? error.message : String(error)}` 
    }
  }
}

// 测试获取最近记录
async function testGetRecent() {
  recentResult.value = null
  try {
    const data = await getRecent(10)
    recentResult.value = { error: false, message: '', data }
  } catch (error) {
    recentResult.value = { 
      error: true, 
      message: `❌ 获取失败: ${error instanceof Error ? error.message : String(error)}` 
    }
  }
}

// 格式化日期
function formatDate(dateStr: string): string {
  if(!dateStr) return dateStr
  const m = dateStr.match(/^(\d{4})-(\d{2})-(\d{2}) (\d{2}):(\d{2}):(\d{2})$/)
  if(m){
    const y = Number(m[1]), mo = Number(m[2]) - 1, d = Number(m[3])
    const hh = Number(m[4]), mm = Number(m[5]), ss = Number(m[6])
    return new Date(y, mo, d, hh, mm, ss).toLocaleString()
  }
  try {
    return new Date(dateStr).toLocaleString()
  } catch {
    return dateStr
  }
}
</script>

<style scoped>
.debugger {
  padding: 16px;
  background: #f5f5f5;
  border-radius: 8px;
  margin: 16px 0;
}

.section {
  margin-bottom: 24px;
  padding: 16px;
  background: white;
  border-radius: 4px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.1);
}

.form-group {
  display: flex;
  gap: 8px;
  margin-bottom: 12px;
  align-items: center;
}

.form-group input {
  padding: 6px 8px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

.btn {
  padding: 6px 12px;
  background: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
}

.btn:hover {
  background: #0056b3;
}

.result {
  margin-top: 12px;
  padding: 8px;
  background: #e8f5e8;
  border-radius: 4px;
  border-left: 4px solid #28a745;
}

.result.error {
  background: #f8e8e8;
  border-left-color: #dc3545;
}

.result ul {
  margin: 8px 0;
  padding-left: 20px;
}

.result li {
  margin: 4px 0;
  font-size: 14px;
}

h3, h4 {
  margin-top: 0;
  color: #333;
}

h3 {
  font-size: 18px;
  margin-bottom: 16px;
}

h4 {
  font-size: 16px;
  margin-bottom: 12px;
}
</style>
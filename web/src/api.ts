export interface SubmitReq { name: string; score: number; comment: string }
export interface RecordVO { id: number; name: string; score: number; comment: string; finished: string }

const API_BASE = (import.meta as any).env?.VITE_API_BASE || 'http://localhost:8888'

export async function submitScore(req: SubmitReq): Promise<{ ok: boolean }>{
  const res = await fetch(`${API_BASE}/api/records/submit`, {
    method: 'POST', headers: { 'Content-Type':'application/json' }, body: JSON.stringify(req)
  })
  if (!res.ok) throw new Error('提交失败')
  return res.json()
}

export async function getRank(limit=50): Promise<RecordVO[]>{
  const res = await fetch(`${API_BASE}/api/records/rank?limit=${limit}`)
  if (!res.ok) throw new Error('获取排行失败')
  return res.json()
}

export async function getRecent(limit=50): Promise<RecordVO[]>{
  const res = await fetch(`${API_BASE}/api/records/recent?limit=${limit}`)
  if (!res.ok) throw new Error('获取最近失败')
  return res.json()
}

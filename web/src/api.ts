export interface SubmitReq { name: string; score: number; comment: string }
export interface RecordVO { id: number; name: string; score: number; comment: string; finished: string }
export type Grade = RecordVO

const API_BASE = (import.meta as any).env?.VITE_API_BASE || ''

export async function submitScore(req: SubmitReq): Promise<{ ok: boolean }>{
  const res = await fetch(`${API_BASE}/api/rank_list/submit`, {
    method: 'POST', headers: { 'Content-Type':'application/json' }, body: JSON.stringify(req)
  })
  if (!res.ok) {
    const txt = await res.text().catch(() => '')
    throw new Error(`提交失败: ${res.status} ${txt}`)
  }
  return res.json()
}

export async function getRank(limit=50): Promise<RecordVO[]>{
  const res = await fetch(`${API_BASE}/api/rank_list/query`, {
    method: 'POST', headers: { 'Content-Type': 'application/json', 'Accept': 'application/json' },
    body: JSON.stringify({ limit, classify: 'sort' })
  })
  if (!res.ok) {
    const txt = await res.text().catch(() => '')
    throw new Error(`获取排行失败: ${res.status} ${txt}`)
  }
  const body = await res.json()
  return body.grade_list || []
}

export async function getRecent(limit=50): Promise<RecordVO[]>{
  const res = await fetch(`${API_BASE}/api/rank_list/query`, {
    method: 'POST', headers: { 'Content-Type': 'application/json', 'Accept': 'application/json' },
    body: JSON.stringify({ limit, classify: 'recent' })
  })
  if (!res.ok) {
    const txt = await res.text().catch(() => '')
    throw new Error(`获取最近失败: ${res.status} ${txt}`)
  }
  const body = await res.json()
  return body.grade_list || []
}

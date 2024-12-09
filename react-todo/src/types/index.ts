export type Task = {
  id: number
  title: string
  created_at: Date
  updated_at: Date
}
export type CsrfToken = {
  csrf_token: string
}
export type Credential = {
  email: string
  password: string
}
export type SignupCredential = {
  email: string
  password: string
  name: string
  age: number
}
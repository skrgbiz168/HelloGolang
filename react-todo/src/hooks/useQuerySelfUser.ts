import axios from 'axios'
import { useQuery } from '@tanstack/react-query'
import { User } from '../types'
import { useError } from '../hooks/useError'

export const useQuerySelfUser = () => {
  const { switchErrorHandling } = useError()
  const getTasks = async () => {
    const { data } = await axios.get<User[]>(
      `${process.env.REACT_APP_API_URL}/selfUser`,
      { withCredentials: true }
    )
    return data
  }
  return useQuery<User[], Error>({
    queryKey: ['Users'],
    queryFn: getTasks,
    staleTime: Infinity,
    onError: (err: any) => {
      if (err.response.data.message) {
        switchErrorHandling(err.response.data.message)
      } else {
        switchErrorHandling(err.response.data)
      }
    },
  })
}
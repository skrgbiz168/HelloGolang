import axios from 'axios'
import { useQuery } from '@tanstack/react-query'
import { User } from '../types'
import { useError } from '../hooks/useError'

export const useQueryUsers = () => {
  const { switchErrorHandling } = useError()
  const getUser = async () => {
    const { data } = await axios.get<User>(
      `${process.env.REACT_APP_API_URL}/selectUsers`,
      { withCredentials: true }
    )
    return data
  }
  const query = useQuery<User, Error>({
    queryKey: ['Users'],
    queryFn: getUser,
    staleTime: Infinity,
    onError: (err: any) => {
      if (err.response.data.message) {
        switchErrorHandling(err.response.data.message)
      } else {
        switchErrorHandling(err.response.data)
      }
    },
  })

  const {data:user, isLoading} = query;
  return { user, isLoading };
}
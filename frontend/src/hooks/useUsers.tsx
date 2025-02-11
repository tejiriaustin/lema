import {useQuery, keepPreviousData, useQueryClient, useMutation} from '@tanstack/react-query'
import {createUser, getUsers} from '../services/api'

export function useUsers(page: number) {
    const queryClient = useQueryClient();

    const data = useQuery({
        queryKey: ['users', page],
        queryFn: () => getUsers(page),
        placeholderData: keepPreviousData,
        staleTime: 5000,
    })

    const createUserMutation = useMutation({
        mutationFn: createUser,
        onSuccess: () => {
            queryClient.invalidateQueries({
                queryKey: ['users']
            }).then(() => {
                console.log("user queries invalidated")
            })
            queryClient.prefetchQuery({
                queryKey: ['users', page],
                queryFn: () => getUsers(page)
            }).then(() => {
                console.log("prefetched users")
            })
        }
    })

    return {
        data,
        createUser: createUserMutation.mutate,
        isLoading: data.isLoading || createUserMutation.isPending,
        error: data.error || createUserMutation.error
    }
}
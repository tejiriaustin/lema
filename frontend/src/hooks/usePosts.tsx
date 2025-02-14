import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { getUserPosts, createPost, deletePost } from '../services/api'

export function usePosts(userId: string) {
    const queryClient = useQueryClient()

    const posts = useQuery({
        queryKey: ['posts', userId],
        queryFn: () => getUserPosts(userId),
        enabled: !!userId
    })

    const createPostMutation = useMutation({
        mutationFn: createPost,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['posts', userId] })
        }
    })

    const deletePostMutation = useMutation({
        mutationFn: deletePost,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['posts', userId] })
        }
    })

    return {
        posts,
        createPost: createPostMutation.mutate,
        deletePost: deletePostMutation.mutate,
        isCreating: createPostMutation.isPending,
        isDeleting: deletePostMutation.isPending,
        isLoading: posts.isLoading || createPostMutation.isPending || deletePostMutation.isPending,
        mutationError: createPostMutation.error,
        isError: createPostMutation.isError,
    }
}
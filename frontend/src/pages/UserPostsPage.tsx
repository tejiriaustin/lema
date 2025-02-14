import { useParams, Link } from "react-router-dom";
import { useState } from "react";
import { usePosts } from "../hooks/usePosts";
import { Loader } from "../components/Loader";
import { PostCard } from "../components/PostCard";
import { CreatePostModal } from "../components/CreatePostModal";
import { DeleteConfirmationModal } from "../components/DeleteConfirmationModal";
import { Post } from '../types';
import { ArrowLeft } from 'lucide-react';
import { NewPostCard } from "../components/NewPostCard.tsx";
import { useToast } from '../components/ToastProvider';

export function UserPostsPage() {
    const { userId } = useParams();
    const [isCreateModalOpen, setIsCreateModalOpen] = useState(false);
    const [postToDelete, setPostToDelete] = useState<string | null>(null);
    const { posts, createPost, deletePost, isLoading: postsLoading } = usePosts(userId!);
    const { showToast } = useToast();

    const handleCreatePost = async (data: { title: string; body: string }) => {
        createPost({ ...data, user_id: userId! }, {
            onSuccess: () => {
                showToast('Post created successfully', 'success');
                setPostToDelete(null);
            },
            onError: (error: Error) => {
                console.log(error?.message);
                showToast('Failed to create post', 'error');
            }
        });
        setIsCreateModalOpen(false);
    };

    const handleDeletePost = async () => {
        if (postToDelete) {
            deletePost(postToDelete, {
                onSuccess: () => {
                    showToast('Post deleted successfully', 'success');
                    setPostToDelete(null);
                },
                onError: (error: Error) => {
                    console.log(error?.message);
                    showToast('Failed to delete post', 'error');
                }
            });
            setPostToDelete(null);
        }
    };

    if (postsLoading) return (
        <div className="flex items-center justify-center">
            <Loader />
        </div>
    );

    return (
        <div className="p-8">
            <div className="max-w-[1000px] mx-auto">
                <div className="text-[16px] mb-6">
                    <Link
                        to="/"
                        className="inline-flex items-center text-gray-600 tracking-wide hover:text-gray-400 transition"
                    >
                        <ArrowLeft className="w-6 h-8 mr-2" />
                        Back to Users
                    </Link>
                </div>

                <div className="mb-8 text-center md:text-left">
                    <h1 className="text-3xl sm:text-4xl md:text-[42px] font-medium mb-2">
                        {posts?.data?.body.user.fullName}
                    </h1>
                    <p className="text-black tracking-wide text-[14px] sm:text-[16px] font-[300]">
                        {posts?.data?.body.user.email} •  <span className="font-[330]"> {posts?.data?.body.posts.length || 0} Posts </span>
                    </p>
                </div>

                <div className="grid gap-6 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 justify-center items-center place-items-center">
                    <NewPostCard onClick={() => setIsCreateModalOpen(true)} />


                    {posts?.data?.body.posts.map((post: Post) => (
                        <PostCard
                            title={post.title}
                            content={post.body}
                            onDelete={() => setPostToDelete(post.id)}
                        />
                    ))}
                </div>

                <CreatePostModal
                    isOpen={isCreateModalOpen}
                    onClose={() => setIsCreateModalOpen(false)}
                    onSubmit={handleCreatePost}
                />

                <DeleteConfirmationModal
                    isOpen={!!postToDelete}
                    onClose={() => setPostToDelete(null)}
                    onConfirm={handleDeletePost}
                />
            </div>
        </div>
    );
}

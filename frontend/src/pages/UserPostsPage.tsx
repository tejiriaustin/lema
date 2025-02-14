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
            <div className="max-w-6xl mx-auto">
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
                    <h1 className="text-2xl sm:text-3xl md:text-4xl font-bold mb-2">
                        {posts?.data?.body.user.fullName}
                    </h1>
                    <p className="text-gray-500 tracking-wide text-[14px] sm:text-[16px]">
                        {posts?.data?.body.user.email} • {posts?.data?.body.posts.length || 0} Posts
                    </p>
                </div>

                {/* ✅ Responsive Grid (2 columns at 1000px) */}
                <div className="grid grid-flow-row grid-cols-1 gap-6 md:grid-cols-2 md:gap-16 lg:grid-cols-3 lg:gap-16 xl:grid-cols-4 lg:gap-16 justify-center items-center place-items-center">
                    <NewPostCard onClick={() => setIsCreateModalOpen(true)} />

                    {posts?.data?.body.posts.map((post: Post) => (
                        <PostCard
                            key={post.id}
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

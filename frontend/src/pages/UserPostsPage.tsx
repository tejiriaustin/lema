import { useParams, Link } from "react-router-dom";
import { useState } from "react";
import { usePosts } from "../hooks/usePosts";
import { Loader } from "../components/Loader";
import { PostCard } from "../components/PostCard";
import { CreatePostModal } from "../components/CreatePostModal";
import { DeleteConfirmationModal } from "../components/DeleteConfirmationModal";
import { Post } from '../types';
import { ArrowLeft } from 'lucide-react';
import {NewPostCard} from "../components/NewPostCard.tsx";

export function UserPostsPage() {
    const { userId } = useParams();
    const [isCreateModalOpen, setIsCreateModalOpen] = useState(false);
    const [postToDelete, setPostToDelete] = useState<string | null>(null);
    const { posts, createPost, deletePost, isLoading: postsLoading } = usePosts(userId!);


    const handleCreatePost = async (data: { title: string; body: string}) => {
        createPost({...data, user_id: userId!});
        setIsCreateModalOpen(false);
    };

    const handleDeletePost = async () => {
        if (postToDelete) {
            deletePost(postToDelete);
            setPostToDelete(null);
        }
    };

    if (postsLoading) return <Loader />;

    return (
        <div className="p-8">
            <div className="max-w-6xl mx-auto">
                <div className="text-[16px]">
                    <Link
                        to="/"
                        className="inline-flex items-center text-gray-600 mb-6 tracking-wide"
                    >
                        <ArrowLeft className="w-6 h-8 mr-2" />
                        Back to Users
                    </Link>
                </div>

                <div className="mb-8">
                    <h1 className="text-4xl mb-2">{posts?.data?.body.user.fullName}</h1>
                    <p className="text-gray-500 tracking-wide text-[16px]">
                        {posts?.data?.body.user.email} • {posts?.data?.body.posts.length || 0} Posts
                    </p>
                </div>

                <div className="grid grid-cols-4 gap-6">
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
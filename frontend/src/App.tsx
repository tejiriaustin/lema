import './App.css';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
import { UsersPage } from './pages/UsersPage'
import { UserPostsPage } from './pages/UserPostsPage'
import {TestPage} from "./pages/test.tsx";
import {ToastProvider} from "./components/ToastProvider.tsx";

const queryClient = new QueryClient({
        defaultOptions: {
            queries: {
                refetchOnWindowFocus: false,
                refetchOnMount: false,
                staleTime: 30000,
                retry: false,
                refetchOnReconnect: false,
            },
        },
    }
);

function App() {
    return (
        <div className="mt-[50px]">
            <QueryClientProvider client={queryClient}>
                <ToastProvider>
                    <BrowserRouter>
                        <Routes>
                            <Route path="/test" element={<TestPage />} />
                            <Route path="/" element={<Navigate to="/users" replace />} />
                            <Route path="/users" element={<UsersPage />} />
                            <Route path="/users/:userId/posts" element={<UserPostsPage />} />
                        </Routes>
                    </BrowserRouter>
                </ToastProvider>
                <ReactQueryDevtools />
            </QueryClientProvider>
        </div>
    )
}

export default App
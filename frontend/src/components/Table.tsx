import { ReactNode } from "react";
import { Loader } from "./Loader";
import { ArrowRight, ArrowLeft } from "lucide-react";

interface Column<T> {
    header: string;
    key: keyof T;
    width?: string;
    render?: (value: T[keyof T], item: T) => ReactNode;
}

interface TableProps<T> {
    data: T[];
    columns: Column<T>[];
    onRowClick?: (item: T) => void;
    isLoading?: boolean;
    rowsPerPage?: number;
    currentPage: number
    onPageChange: (page: number) => void
    totalPages: number
}

export function Table<T>({data, columns, onRowClick, isLoading, totalPages, onPageChange, currentPage}: TableProps<T>) {
    const goToPage = (page: number) => {
        if (page >= 1 && page <= totalPages) {
            onPageChange(page);
        }
    };

    const renderPaginationNumbers = () => {
        const pages = [];

        if (currentPage <= 3) {
            pages.push(1, 2, 3);
            if (totalPages > 6) {
                pages.push('...');
                pages.push(totalPages - 2, totalPages - 1, totalPages);
            } else {
                for (let i = 4; i <= totalPages; i++) {
                    pages.push(i);
                }
            }
        }
        else {
            pages.push(currentPage - 1, currentPage, currentPage + 1);

            if (currentPage < totalPages - 3) {
                pages.push('...');
                pages.push(totalPages - 2, totalPages - 1, totalPages);
            } else {
                for (let i = currentPage + 2; i <= totalPages; i++) {
                    pages.push(i);
                }
            }
        }

        return pages;
    };

    return (
        <div className="flex flex-col">
            <div className="overflow-x-auto">
                <div className="rounded-lg shadow-sm overflow-hidden border border-gray-200 font-inter min-w-[640px]">
                    <table className="w-full border-collapse">
                        <thead>
                        <tr>
                            {columns.map((column) => (
                                <th
                                    key={String(column.key)}
                                    className="text-left text-[14px] font-normal text-[#535862] py-5 px-6"
                                    style={{width: column.width}}
                                >
                                    {column.header}
                                </th>
                            ))}
                        </tr>
                        </thead>
                        <tbody>
                        {isLoading ? (
                            <tr>
                                <td colSpan={columns.length} className="h-64">
                                    <div className="flex items-center justify-center h-full">
                                        <Loader/>
                                    </div>
                                </td>
                            </tr>
                        ) : data.length > 0 ? (
                            data.map((item, index) => (
                                <tr
                                    key={index}
                                    onClick={() => onRowClick?.(item)}
                                    className="cursor-pointer border-b last:border-none hover:bg-gray-50"
                                >
                                    {columns.map((column) => (
                                        <td key={String(column.key)}
                                            className={`text-[16px] py-7 px-6 tracking-wide text-[#535862] whitespace-nowrap ${
                                                column.key === 'address' ? 'truncate max-w-[1px]' : ''
                                            } ${
                                                column.key === 'fullName'
                                                    ? 'font-normal'
                                                    : 'font-light'
                                            }`}>
                                            {column.render
                                                ? column.render(item[column.key], item)
                                                : String(item[column.key] || "")}
                                        </td>
                                    ))}
                                </tr>
                            ))
                        ) : (
                            <tr>
                                <td colSpan={columns.length} className="text-center text-gray-500 py-4">
                                    No data available.
                                </td>
                            </tr>
                        )}
                        </tbody>
                    </table>
                </div>
            </div>

            <div className="flex justify-end mt-4 px-2 py-4 bg-white text-gray-400">
                <div className="flex items-center gap-2">
                    <button
                        className={`flex items-center gap-3 text-[16px] mr-2 font-[475] text-gray-800 hover:text-gray-900`}
                        onClick={() => goToPage(currentPage - 1)}
                        disabled={currentPage === 1}
                    >
                        <ArrowLeft size={18}/>
                        Previous
                    </button>

                    <div className="flex gap-0.5 px-8 text-[16px] font-[425]">
                        {renderPaginationNumbers().map((page, index) => (
                            <button
                                key={index}
                                className={`w-12 h-12 flex items-center justify-center rounded-lg ${
                                    page === '...'
                                        ? 'cursor-default'
                                        : page === currentPage
                                            ? "bg-purple-50 text-[#7F56D9]"
                                            : "text-gray-700 hover:bg-gray-100"
                                }`}
                                onClick={() => typeof page === 'number' && goToPage(page)}
                                disabled={page === '...'}
                            >
                                {page}
                            </button>
                        ))}
                    </div>

                    <button
                        className={`flex items-center gap-2 text-[14.3px] font-[550] ${
                            currentPage === totalPages ? "text-gray-400 cursor-not-allowed" : "text-gray-700 hover:text-gray-900"
                        }`}
                        onClick={() => goToPage(currentPage + 1)}
                        disabled={currentPage === totalPages}
                    >
                        Next
                        <ArrowRight size={18}/>
                    </button>
                </div>
            </div>
        </div>
    );
}

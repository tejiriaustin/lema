import {Table} from "../components/Table.tsx";
import {Button} from "../components/Button.tsx";
import {PostForm} from "../components/PostForm.tsx";
import {Loader} from "../components/Loader.tsx";
import {PostCard} from "../components/PostCard.tsx";
import {NewPostCard} from "../components/NewPostCard.tsx";

export function TestPage() {

    const testUsers = [
        { name: "James Sunderland", email: "james.sunderland@acme.corp", address: "11 Katz St., Pennsylvania, Centralia, M4A2T6" },
        { name: "Heather Mayson", email: "h.mayson@acme.corp", address: "24 Lindsey St., British Columbia, Vancouver, N9M2..." },
        { name: "Henry Townshend", email: "henry_townsend@acme.corp", address: "10 Rendell St., Ontario, Toronto, M2K3B8" },
        { name: "Walter Sullivan", email: "walter.s@acme.corp", address: "9 Wiltse Road, Alberta, Canmore, N9W4H9" },
        { name: "Maria Garland", email: "maria.garland@acme.corp", address: "21 Garnet St., Manitoba, Winnipeg, M1N4P2" },
        { name: "Alex Shepherd", email: "alex.shepherd@acme.corp", address: "16 Maple Rd., Quebec, Montreal, L3B5H7" },
        { name: "James Sunderland", email: "james.sunderland@acme.corp", address: "11 Katz St., Pennsylvania, Centralia, M4A2T6" },
        { name: "Heather Mayson", email: "h.mayson@acme.corp", address: "24 Lindsey St., British Columbia, Vancouver, N9M2..." },
        { name: "Henry Townshend", email: "henry_townsend@acme.corp", address: "10 Rendell St., Ontario, Toronto, M2K3B8" },
        { name: "Walter Sullivan", email: "walter.s@acme.corp", address: "9 Wiltse Road, Alberta, Canmore, N9W4H9" },
        { name: "Maria Garland", email: "maria.garland@acme.corp", address: "21 Garnet St., Manitoba, Winnipeg, M1N4P2" },
        { name: "Alex Shepherd", email: "alex.shepherd@acme.corp", address: "16 Maple Rd., Quebec, Montreal, L3B5H7" },
        { name: "James Sunderland", email: "james.sunderland@acme.corp", address: "11 Katz St., Pennsylvania, Centralia, M4A2T6" },
        { name: "Heather Mayson", email: "h.mayson@acme.corp", address: "24 Lindsey St., British Columbia, Vancouver, N9M2..." },
        { name: "Henry Townshend", email: "henry_townsend@acme.corp", address: "10 Rendell St., Ontario, Toronto, M2K3B8" },
        { name: "Walter Sullivan", email: "walter.s@acme.corp", address: "9 Wiltse Road, Alberta, Canmore, N9W4H9" },
        { name: "Maria Garland", email: "maria.garland@acme.corp", address: "21 Garnet St., Manitoba, Winnipeg, M1N4P2" },
        { name: "Alex Shepherd", email: "alex.shepherd@acme.corp", address: "16 Maple Rd., Quebec, Montreal, L3B5H7" },
        { name: "James Sunderland", email: "james.sunderland@acme.corp", address: "11 Katz St., Pennsylvania, Centralia, M4A2T6" },
        { name: "Heather Mayson", email: "h.mayson@acme.corp", address: "24 Lindsey St., British Columbia, Vancouver, N9M2..." },
        { name: "Henry Townshend", email: "henry_townsend@acme.corp", address: "10 Rendell St., Ontario, Toronto, M2K3B8" },
        { name: "Walter Sullivan", email: "walter.s@acme.corp", address: "9 Wiltse Road, Alberta, Canmore, N9W4H9" },
        { name: "Maria Garland", email: "maria.garland@acme.corp", address: "21 Garnet St., Manitoba, Winnipeg, M1N4P2" },
        { name: "Alex Shepherd", email: "alex.shepherd@acme.corp", address: "16 Maple Rd., Quebec, Montreal, L3B5H7" },
        { name: "James Sunderland", email: "james.sunderland@acme.corp", address: "11 Katz St., Pennsylvania, Centralia, M4A2T6" },
        { name: "Heather Mayson", email: "h.mayson@acme.corp", address: "24 Lindsey St., British Columbia, Vancouver, N9M2..." },
        { name: "Henry Townshend", email: "henry_townsend@acme.corp", address: "10 Rendell St., Ontario, Toronto, M2K3B8" },
        { name: "Walter Sullivan", email: "walter.s@acme.corp", address: "9 Wiltse Road, Alberta, Canmore, N9W4H9" },
        { name: "Maria Garland", email: "maria.garland@acme.corp", address: "21 Garnet St., Manitoba, Winnipeg, M1N4P2" },
        { name: "Alex Shepherd", email: "alex.shepherd@acme.corp", address: "16 Maple Rd., Quebec, Montreal, L3B5H7" },
        { name: "James Sunderland", email: "james.sunderland@acme.corp", address: "11 Katz St., Pennsylvania, Centralia, M4A2T6" },
        { name: "Heather Mayson", email: "h.mayson@acme.corp", address: "24 Lindsey St., British Columbia, Vancouver, N9M2..." },
        { name: "Henry Townshend", email: "henry_townsend@acme.corp", address: "10 Rendell St., Ontario, Toronto, M2K3B8" },
        { name: "Walter Sullivan", email: "walter.s@acme.corp", address: "9 Wiltse Road, Alberta, Canmore, N9W4H9" },
        { name: "Maria Garland", email: "maria.garland@acme.corp", address: "21 Garnet St., Manitoba, Winnipeg, M1N4P2" },
        { name: "Alex Shepherd", email: "alex.shepherd@acme.corp", address: "16 Maple Rd., Quebec, Montreal, L3B5H7" },
        { name: "James Sunderland", email: "james.sunderland@acme.corp", address: "11 Katz St., Pennsylvania, Centralia, M4A2T6" },
        { name: "Heather Mayson", email: "h.mayson@acme.corp", address: "24 Lindsey St., British Columbia, Vancouver, N9M2..." },
        { name: "Henry Townshend", email: "henry_townsend@acme.corp", address: "10 Rendell St., Ontario, Toronto, M2K3B8" },
        { name: "Walter Sullivan", email: "walter.s@acme.corp", address: "9 Wiltse Road, Alberta, Canmore, N9W4H9" },
        { name: "Maria Garland", email: "maria.garland@acme.corp", address: "21 Garnet St., Manitoba, Winnipeg, M1N4P2" },
        { name: "Alex Shepherd", email: "alex.shepherd@acme.corp", address: "16 Maple Rd., Quebec, Montreal, L3B5H7" },
        { name: "James Sunderland", email: "james.sunderland@acme.corp", address: "11 Katz St., Pennsylvania, Centralia, M4A2T6" },
        { name: "Heather Mayson", email: "h.mayson@acme.corp", address: "24 Lindsey St., British Columbia, Vancouver, N9M2..." },
        { name: "Henry Townshend", email: "henry_townsend@acme.corp", address: "10 Rendell St., Ontario, Toronto, M2K3B8" },
        { name: "Walter Sullivan", email: "walter.s@acme.corp", address: "9 Wiltse Road, Alberta, Canmore, N9W4H9" },
        { name: "Maria Garland", email: "maria.garland@acme.corp", address: "21 Garnet St., Manitoba, Winnipeg, M1N4P2" },
        { name: "Alex Shepherd", email: "alex.shepherd@acme.corp", address: "16 Maple Rd., Quebec, Montreal, L3B5H7" },
        { name: "James Sunderland", email: "james.sunderland@acme.corp", address: "11 Katz St., Pennsylvania, Centralia, M4A2T6" },
        { name: "Heather Mayson", email: "h.mayson@acme.corp", address: "24 Lindsey St., British Columbia, Vancouver, N9M2..." },
        { name: "Henry Townshend", email: "henry_townsend@acme.corp", address: "10 Rendell St., Ontario, Toronto, M2K3B8" },
        { name: "Walter Sullivan", email: "walter.s@acme.corp", address: "9 Wiltse Road, Alberta, Canmore, N9W4H9" },
        { name: "Maria Garland", email: "maria.garland@acme.corp", address: "21 Garnet St., Manitoba, Winnipeg, M1N4P2" },
        { name: "Alex Shepherd", email: "alex.shepherd@acme.corp", address: "16 Maple Rd., Quebec, Montreal, L3B5H7" },
        { name: "James Sunderland", email: "james.sunderland@acme.corp", address: "11 Katz St., Pennsylvania, Centralia, M4A2T6" },
        { name: "Heather Mayson", email: "h.mayson@acme.corp", address: "24 Lindsey St., British Columbia, Vancouver, N9M2..." },
        { name: "Henry Townshend", email: "henry_townsend@acme.corp", address: "10 Rendell St., Ontario, Toronto, M2K3B8" },
        { name: "Walter Sullivan", email: "walter.s@acme.corp", address: "9 Wiltse Road, Alberta, Canmore, N9W4H9" },
        { name: "Maria Garland", email: "maria.garland@acme.corp", address: "21 Garnet St., Manitoba, Winnipeg, M1N4P2" },
        { name: "Alex Shepherd", email: "alex.shepherd@acme.corp", address: "16 Maple Rd., Quebec, Montreal, L3B5H7" },
        { name: "James Sunderland", email: "james.sunderland@acme.corp", address: "11 Katz St., Pennsylvania, Centralia, M4A2T6" },
        { name: "Heather Mayson", email: "h.mayson@acme.corp", address: "24 Lindsey St., British Columbia, Vancouver, N9M2..." },
        { name: "Henry Townshend", email: "henry_townsend@acme.corp", address: "10 Rendell St., Ontario, Toronto, M2K3B8" },
        { name: "Walter Sullivan", email: "walter.s@acme.corp", address: "9 Wiltse Road, Alberta, Canmore, N9W4H9" },
        { name: "Maria Garland", email: "maria.garland@acme.corp", address: "21 Garnet St., Manitoba, Winnipeg, M1N4P2" },
        { name: "Alex Shepherd", email: "alex.shepherd@acme.corp", address: "16 Maple Rd., Quebec, Montreal, L3B5H7" },
        { name: "James Sunderland", email: "james.sunderland@acme.corp", address: "11 Katz St., Pennsylvania, Centralia, M4A2T6" },
        { name: "Heather Mayson", email: "h.mayson@acme.corp", address: "24 Lindsey St., British Columbia, Vancouver, N9M2..." },
        { name: "Henry Townshend", email: "henry_townsend@acme.corp", address: "10 Rendell St., Ontario, Toronto, M2K3B8" },
        { name: "Walter Sullivan", email: "walter.s@acme.corp", address: "9 Wiltse Road, Alberta, Canmore, N9W4H9" },
        { name: "Maria Garland", email: "maria.garland@acme.corp", address: "21 Garnet St., Manitoba, Winnipeg, M1N4P2" },
        { name: "Alex Shepherd", email: "alex.shepherd@acme.corp", address: "16 Maple Rd., Quebec, Montreal, L3B5H7" },
        { name: "James Sunderland", email: "james.sunderland@acme.corp", address: "11 Katz St., Pennsylvania, Centralia, M4A2T6" },
        { name: "Heather Mayson", email: "h.mayson@acme.corp", address: "24 Lindsey St., British Columbia, Vancouver, N9M2..." },
        { name: "Henry Townshend", email: "henry_townsend@acme.corp", address: "10 Rendell St., Ontario, Toronto, M2K3B8" },
        { name: "Walter Sullivan", email: "walter.s@acme.corp", address: "9 Wiltse Road, Alberta, Canmore, N9W4H9" },
        { name: "Maria Garland", email: "maria.garland@acme.corp", address: "21 Garnet St., Manitoba, Winnipeg, M1N4P2" },
        { name: "Alex Shepherd", email: "alex.shepherd@acme.corp", address: "16 Maple Rd., Quebec, Montreal, L3B5H7" },
        { name: "James Sunderland", email: "james.sunderland@acme.corp", address: "11 Katz St., Pennsylvania, Centralia, M4A2T6" },
        { name: "Heather Mayson", email: "h.mayson@acme.corp", address: "24 Lindsey St., British Columbia, Vancouver, N9M2..." },
        { name: "Henry Townshend", email: "henry_townsend@acme.corp", address: "10 Rendell St., Ontario, Toronto, M2K3B8" },
        { name: "Walter Sullivan", email: "walter.s@acme.corp", address: "9 Wiltse Road, Alberta, Canmore, N9W4H9" },
        { name: "Maria Garland", email: "maria.garland@acme.corp", address: "21 Garnet St., Manitoba, Winnipeg, M1N4P2" },
        { name: "Alex Shepherd", email: "alex.shepherd@acme.corp", address: "16 Maple Rd., Quebec, Montreal, L3B5H7" },
        { name: "James Sunderland", email: "james.sunderland@acme.corp", address: "11 Katz St., Pennsylvania, Centralia, M4A2T6" },
        { name: "Heather Mayson", email: "h.mayson@acme.corp", address: "24 Lindsey St., British Columbia, Vancouver, N9M2..." },
        { name: "Henry Townshend", email: "henry_townsend@acme.corp", address: "10 Rendell St., Ontario, Toronto, M2K3B8" },
        { name: "Walter Sullivan", email: "walter.s@acme.corp", address: "9 Wiltse Road, Alberta, Canmore, N9W4H9" },
        { name: "Maria Garland", email: "maria.garland@acme.corp", address: "21 Garnet St., Manitoba, Winnipeg, M1N4P2" },
        { name: "Alex Shepherd", email: "alex.shepherd@acme.corp", address: "16 Maple Rd., Quebec, Montreal, L3B5H7" },
        { name: "James Sunderland", email: "james.sunderland@acme.corp", address: "11 Katz St., Pennsylvania, Centralia, M4A2T6" },
        { name: "Heather Mayson", email: "h.mayson@acme.corp", address: "24 Lindsey St., British Columbia, Vancouver, N9M2..." },
        { name: "Henry Townshend", email: "henry_townsend@acme.corp", address: "10 Rendell St., Ontario, Toronto, M2K3B8" },
        { name: "Walter Sullivan", email: "walter.s@acme.corp", address: "9 Wiltse Road, Alberta, Canmore, N9W4H9" },
        { name: "Maria Garland", email: "maria.garland@acme.corp", address: "21 Garnet St., Manitoba, Winnipeg, M1N4P2" },
        { name: "Alex Shepherd", email: "alex.shepherd@acme.corp", address: "16 Maple Rd., Quebec, Montreal, L3B5H7" },
    ];

    const columns = [
        { header: "Full Name", key: "name" as const, width: "200px" },
        { header: "Email Address", key: "email" as const, width: "300px" },
        { header: "Address", key: "address" as const, width: "392px" },
    ];


    return (
            <div className="p-8 font-inter"> {/* Added font-inter */}
                <h1 className="text-5xl mb-6">Users</h1>
                <div className="rounded-lg shadow-sm w-[856px]">
                    <Table data={testUsers} columns={columns} isLoading={false} rowsPerPage={4} currentPage={1} onPageChange={()=>{}} totalPages={0}/>
                </div>

                <div className="flex gap-3">
                    <Button title="Cancel"/>
                    <Button title="Publish" variant="primary"/>
                </div>

                <PostForm/>
                <Loader />

                <PostCard
                    title="What a Nice Town"
                    content="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dol. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dol"
                    onDelete={() => console.log('delete post')}
                />
                <NewPostCard onClick={() => {
                    console.log('Create new post');
                }} />
            </div>
    );
}
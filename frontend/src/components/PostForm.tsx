import {useState} from "react";
import {InputBox} from "./InputBox.tsx";
import {Button} from "./Button.tsx";

export function PostForm() {
    const [title, setTitle] = useState("");
    const [content, setContent] = useState("");

    return (
        <div className="space-y-6">
            <InputBox
                label="Post title"
                placeholder="Give your post a title"
                value={title}
                onChange={setTitle}
            />

            <InputBox
                label="Post content"
                placeholder="Write something mind-blowing"
                value={content}
                onChange={setContent}
                multiline
            />

            <div className="flex gap-3 justify-end">
                <Button title="Cancel" />
                <Button
                    title="Publish"
                    variant="primary"
                    disabled={!title || !content}
                />
            </div>
        </div>
    );
}
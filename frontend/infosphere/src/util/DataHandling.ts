import { FormEvent } from "react";

export function extractDataFromFormEvent(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const fd = new FormData(event.currentTarget);
    return Object.fromEntries(fd.entries());
}
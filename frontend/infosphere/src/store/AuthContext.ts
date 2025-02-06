/* 
User context should expose:
- User Model
- Login
- Logout
- Signup
*/
import React, { createContext, PropsWithChildren, FormEvent, useState } from "react";
import User from "../models/User";
import { extractDataFromFormEvent } from "../util/DataHandling";
import { useMutation, useQuery } from '@tanstack/react-query';
import axios from 'axios'

interface AuthContextType {
    user: User | null;
    login: () => void
    logout: () => void
    signup: () => void
}

export const AuthContext = createContext<AuthContextType>({
    user: null,
    login: () => {},
    logout: () => {},
    signup: () => {},
});

interface UserSignupDetails {
    username: string
    email: string
    // password: string
}

async function createUser(details: UserSignupDetails) {
    const query = useMutation({
        mutationFn: () => createUserRequest()
    })
}

function createUserRequest(details: UserSignupDetails) {
   
}

export default function AuthContextProvider = ({children}: PropsWithChildren): React.JSX.Element => {

    const [user, setUser] = useState<User | null>(null);

    // function login(event: FormEvent<HTMLFormElement>) {
    //     const data = extractDataFromFormEvent(event);

    //     event.currentTarget.reset();
    // }

    function signup(event: FormEvent<HTMLFormElement>) {
        const data = extractDataFromFormEvent(event);

    }

    return ( 
        <AuthContext.Provider value=>
        { children }
        </AuthContext> 
    );
}
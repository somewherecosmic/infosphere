import { createBrowserRouter} from 'react-router';
import Login from './components/Login';
import Signup from './components/Signup';



export default router = createBrowserRouter([
    { 
        path: "/login",
        element: <Login />
    },
    {
        path: "/signup",
        element: <Signup />
    }
])
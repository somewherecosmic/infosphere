
import { Routes, Route } from 'react-router'
import './App.css'
import Landing from './components/Landing'
import Signup from './components/Signup'


const App = () => {

  return (
    <Routes>
      <Route path="/" element={<Landing />} />
      <Route path="/signup" element={<Signup />} />
    </Routes>
  )
}

export default App

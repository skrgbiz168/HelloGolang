import { useEffect } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { Auth } from './components/Auth'
import { DashBoard } from './components/DashBoard'
import { Todo } from './components/tasks/Todo'
import { Select } from './components/users/Select'
import axios from 'axios';
import { CsrfToken } from './types';

function App() {
  useEffect(() => {
    axios.defaults.withCredentials = true
    const getCsrfToken = async () => {
      const { data } = await axios.get<CsrfToken>(
        `${process.env.REACT_APP_API_URL}/csrf`
      )
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
    }
    getCsrfToken()
  }, [])

  return (
    <BrowserRouter>
      <Routes>
        <Route path="" element={<Auth />} />
        <Route path="/dashboard" element={<DashBoard />} />
        <Route path="/todo" element={<Todo />} />
        <Route path="/userSelect" element={<Select />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;

import React from 'react';
import Auth from "./pages/auth/Auth";
import {BrowserRouter, Route, Routes} from "react-router-dom";
import Home from "./pages/home/Home";
import Info from "./pages/info/Info";

const App: React.FC = () => {
    return (
        <BrowserRouter>
            <div className="App">
                <Routes>
                    <Route path="/" element={<Auth/>}/>
                    <Route path="/home" element={<Home/>}/>
                    <Route path="/info" element={<Info/>}/>
                </Routes>
            </div>
            </BrowserRouter>
    );
};

export default App;
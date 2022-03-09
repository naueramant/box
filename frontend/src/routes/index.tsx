import { FunctionComponent, Suspense } from "react";
import { Route, BrowserRouter as Router, Routes } from "react-router-dom";
import FilesView from "views/Files";
import LoginView from "views/Login";

const AppRoutes: FunctionComponent = () => {
    return (
        <Router>
            <Suspense fallback="">
                <Routes>
                    <Route path="/" element={<FilesView />} />
                    <Route path="/login" element={<LoginView />} />
                </Routes>
            </Suspense>
        </Router>
    );
};

export default AppRoutes;

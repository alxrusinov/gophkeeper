import { Route, Routes } from "react-router-dom";
import { MainPage } from "./pages/main-page";

function App() {
  return (
    <Routes>
      <Route path="/">
        <MainPage />
      </Route>
    </Routes>
  );
}

export default App;

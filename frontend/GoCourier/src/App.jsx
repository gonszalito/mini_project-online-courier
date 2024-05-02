import { RouterProvider, createBrowserRouter } from "react-router-dom";
import {
  HomeLayout,
  Landing,
  LandingCourier,
  Login,
  Register,
  Profile,
  OrderHistory,
  CourierOrderHistory,
  CreateOrder
} from "./pages";
import { ToastContainer } from "react-toastify";

const router = createBrowserRouter([
  {
    path: "/",
    element: <HomeLayout />,
    children: [
      {
        index: true,
        element: <LandingContainer />,
      },
      {
        path: "login",
        element: <Login />,
      },
      {
        path: "register",
        element: <Register />,
      },
      {
        path: "user-profile",
        element: <Profile />,
      },
      {
        path:"order-history",
        element: <OrderHistoryContainer />, 
      },
      {
        path:"new-order",
        element: <CreateOrder />
      }
    ],
  },
]);

function LandingContainer() {
  const role = localStorage.getItem("role");

  return (
    <>
      {role === "user" && <Landing />}
      {role === "courier" && <LandingCourier />}
    </>
  );
}

function OrderHistoryContainer() {
  const role = localStorage.getItem("role");


  return (
    <>
      {role === "user" && <OrderHistory />}
      {role === "courier" && <CourierOrderHistory />}
    </>
  );
}

function App() {
  return (
    <>
      <RouterProvider router={router} />
      <ToastContainer position="top-center" />
    </>
  );
}

export default App;

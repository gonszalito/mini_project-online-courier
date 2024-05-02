import React, { useState, useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";
import axios from "axios";

const Landing = () => {
  const [orders, setOrders] = useState([]);
  const navigate = useNavigate();

  // Fetch orders when component mounts
  useEffect(() => {
    fetchOrders();
  }, []);

  // Fetch orders periodically
  useEffect(() => {
    const interval = setInterval(fetchOrders, 60000); // Fetch orders every minute
    return () => clearInterval(interval); // Clean up the interval on component unmount
  }, []);

  const fetchOrders = async () => {
    // Get the token from local storage
    const token = localStorage.getItem("token");

    try {
      const response = await axios.get(`http://localhost:8020/order/user`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      const data = response.data;

      if (!data || !data.data) {
        setOrders([]);
      } else {
        // Filter orders with status "pending" or "ongoing"
        const filteredOrders = data.data.filter(order => order.status === "pending" || order.status === "ongoing");
        setOrders(filteredOrders);
      }
    } catch (error) {
      console.error("Error loading landing data:", error);
      setOrders([]);
    }
  };

  const handleDeleteOrder = async (orderId) => {
    try {
      console.log("Deleting order with ID:", orderId);
      // Make a PATCH request to update the order status to "cancelled"
      const token = localStorage.getItem("token");
      const response = await axios.patch(
        `http://localhost:8020/order/update`,
        {
          _id: orderId,
          status: "cancelled",
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      console.log("Delete order response:", response.data);

      // Re-fetch orders after deleting
      fetchOrders();
    } catch (error) {
      console.error("Error deleting order:", error);
    }
  };

  return (
    <main>
      <div className="selected-products">
        <h1 className="text-6xl text-center my-12 max-md:text-4xl text-accent-content">
          Ongoing Orders
        </h1>
        <div className="selected-products-grid max-w-7xl mx-auto flex flex-col items-center">
          <Link
            to="/new-order"
            className="btn btn-wide bg-blue-600 hover:bg-blue-500 text-white"
            style={{ alignSelf: "center", marginBottom: "20px" }}
          >
            New Order
          </Link>
          <div className="orders-tabs">
            {orders.length === 0 ? (
              <p className="text-center text-gray-500 my-8">You have no ongoing orders.</p>
            ) : (
              orders.map((order) => (
                <div key={order._id} className="order-tab" style={{ marginBottom: "20px" }}>
                  <div className="order-content">
                    <p><strong>Description:</strong> {order.description}</p>
                    <p><strong>Start Location:</strong> {order.start_location}</p>
                    <p><strong>End Location:</strong> {order.end_location}</p>
                    <p><strong>Start Time:</strong> {new Date(order.start_time).toLocaleString()}</p>
                    <p><strong>Status:</strong> {order.status}</p>
                  </div>
                  {order.status === "pending" && (
                    <button
                      className="btn  bg-red-600 hover:bg-red-500 text-white mx-auto flex flex-col items-center"
                      onClick={() => handleDeleteOrder(order._id)}
                    >
                      Delete
                    </button>
                  )}
                </div>
              ))
            )}
          </div>
        </div>
      </div>
    </main>
  );
};

export default Landing;

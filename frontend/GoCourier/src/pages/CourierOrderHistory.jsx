import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

const CourierOrderHistory = () => {
  const [orders, setOrders] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    fetchOrders();
  }, []); // Run only once when the component mounts

  const fetchOrders = async () => {
    // Get the token from local storage
    const token = localStorage.getItem("token");

    try {
      const response = await axios.get(`http://localhost:8020/order/courier`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      const data = response.data;
      console.log("Response data:", data); // Log the response data

      // Check if data is null or data.data is null, return an empty array
      if (!data || !data.data) {
        setOrders([]);
      } else {
        setOrders(data.data);
      }
    } catch (error) {
      // Handle error
      console.error("Error loading order data:", error);
      setOrders([]); // Set orders to an empty array in case of error
    }
  };

  // Filter orders with status "finished" or "cancelled"
  const filteredOrders = orders ? orders.filter((order) => order.status === "finished" ) : [];

  return (
    <main>
      <div className="selected-products">
        <h1 className="text-6xl text-center my-12 max-md:text-4xl text-accent-content">
          Past Orders
        </h1>
        <div className="selected-products-grid max-w-7xl mx-auto flex flex-col items-center">
          <div className="orders-tabs">
            {filteredOrders.length === 0 ? (
              <p className="text-center text-gray-500 my-8">You have no past orders.</p>
            ) : (
              filteredOrders.map((order) => (
                <div key={order._id} className="order-tab" style={{ marginBottom: "20px" }}> {/* Add marginBottom for spacing */}
                  <div className="order-content">
                    <p><strong>Description:</strong> {order.description}</p>
                    <p><strong>Start Location:</strong> {order.start_location}</p>
                    <p><strong>End Location:</strong> {order.end_location}</p>
                    <p><strong>Start Time:</strong> {new Date(order.start_time).toLocaleString()}</p>
                    <p><strong>Status:</strong> {order.status}</p>
                  </div>
                </div>
              ))
            )}
          </div>
        </div>
      </div>
    </main>
  );
};

export default CourierOrderHistory;

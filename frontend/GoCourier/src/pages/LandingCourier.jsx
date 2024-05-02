import React, { useState, useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";
import axios from "axios";

const LandingCourier = () => {
  const [pendingOrders, setPendingOrders] = useState([]);
  const [ongoingOrders, setOngoingOrders] = useState([]);
  const navigate = useNavigate();

  // Fetch pending orders when component mounts
  useEffect(() => {
    fetchOngoingOrders();
    fetchPendingOrders();
  }, []);

  // Fetch ongoing orders periodically
  useEffect(() => {
    const interval = setInterval(fetchOngoingOrders, 60000); // Fetch ongoing orders every minute
    return () => clearInterval(interval); // Clean up the interval on component unmount
  }, []);

  const fetchPendingOrders = async () => {
    // Get the token from local storage
    const token = localStorage.getItem("token");

    try {
      const response = await axios.get(`http://localhost:8020/order/pending`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      const data = response.data;

      if (!data || !data.data) {
        setPendingOrders([]);
      } else {
        // Filter orders
        const pending = data.data.filter(order => order.status === "pending");
        setPendingOrders(pending);
      }
    } catch (error) {
      console.error("Error loading pending orders:", error);
      setPendingOrders([]);
    }
  };

  const fetchOngoingOrders = async () => {
    // Get the token from local storage
    const token = localStorage.getItem("token");

    try {
      const response = await axios.get(`http://localhost:8020/order/courier`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });


      const data = response.data;



      if (!data || !data.data) {
        setOngoingOrders([]);
      } else {
        // Filter orders
        const ongoing = data.data.filter(order => order.status === "ongoing");
        setOngoingOrders(ongoing);
      }
    } catch (error) {
      console.error("Error loading ongoing orders:", error);
      setOngoingOrders([]);
    }
  };

  const handleAcceptOrder = async (orderId) => {
    try {
      // Make a PATCH request to update the order status to "ongoing"
      const token = localStorage.getItem("token");
      await axios.patch(
        `http://localhost:8020/order/update`,
        {
          _id: orderId,
          role: "courier",
          status: "ongoing",
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );
  
      // Update local state to remove the accepted order from pending orders
      setPendingOrders(prevOrders => prevOrders.filter(order => order._id !== orderId));
  
      // Fetch ongoing orders to update the list
      fetchOngoingOrders();
    } catch (error) {
      console.error("Error accepting order:", error);
    }
  };
  
  const handleFinishOrder = async (orderId) => {
    try {
      // Make a PATCH request to update the order status to "finished"
      const token = localStorage.getItem("token");
      await axios.patch(
        `http://localhost:8020/order/update`,
        {
          _id: orderId,
          role : "courier",
          status: "finished",
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );
  
      // Update local state to remove the finished order from ongoing orders
      setOngoingOrders(prevOrders => prevOrders.filter(order => order._id !== orderId));
  
      // Fetch pending orders to update the list
      fetchPendingOrders();
    } catch (error) {
      console.error("Error finishing order:", error);
    }
  };
  
  return (
    <main>
      <div className="selected-products">
        <h1 className="text-6xl text-center my-12 max-md:text-4xl text-accent-content">
          Pending Orders
        </h1>
        <div className="selected-products-grid max-w-7xl mx-auto flex flex-col items-center">
          <div className="orders-tabs">
            {pendingOrders.length === 0 ? (
              <p className="text-center text-gray-500 my-8">You have no pending orders.</p>
            ) : (
              pendingOrders.map((order) => (
                <div key={order._id} className="order-tab" style={{ marginBottom: "20px" }}>
                  <div className="order-content">
                    <p><strong>Description:</strong> {order.description}</p>
                    <p><strong>Start Location:</strong> {order.start_location}</p>
                    <p><strong>End Location:</strong> {order.end_location}</p>
                    <p><strong>Start Time:</strong> {new Date(order.start_time).toLocaleString()}</p>
                    <p><strong>Status:</strong> {order.status}</p>
                  </div>
                  <button
                    className="btn bg-green-600 hover:bg-green-500 text-white mx-auto flex flex-col items-center"
                    onClick={() => handleAcceptOrder(order._id)}
                  >
                    Accept Order
                  </button>
                </div>
              ))
            )}
          </div>
        </div>
      </div>
      
      <div className="selected-products">
        <h1 className="text-6xl text-center my-12 max-md:text-4xl text-accent-content">
          Ongoing Orders
        </h1>
        <div className="selected-products-grid max-w-7xl mx-auto flex flex-col items-center">
          <div className="orders-tabs">
            {ongoingOrders.length === 0 ? (
              <p className="text-center text-gray-500 my-8">You have no ongoing orders.</p>
            ) : (
              ongoingOrders.map((order) => (
                <div key={order._id} className="order-tab" style={{ marginBottom: "20px" }}>
                  <div className="order-content">
                    <p><strong>Description:</strong> {order.description}</p>
                    <p><strong>Start Location:</strong> {order.start_location}</p>
                    <p><strong>End Location:</strong> {order.end_location}</p>
                    <p><strong>Start Time:</strong> {new Date(order.start_time).toLocaleString()}</p>
                    <p><strong>Status:</strong> {order.status}</p>
                  </div>
                  <button
                    className="btn bg-green-600 hover:bg-green-500 text-white mx-auto flex flex-col items-center"
                    onClick={() => handleFinishOrder(order._id)}
                  >
                    Finish Order
                  </button>
                </div>
              ))
            )}
          </div>
        </div>
      </div>
    </main>
  );
};

export default LandingCourier;

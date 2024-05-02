import React, { useState } from "react";
import axios from "axios";
import { toast } from "react-toastify";
import { useSelector } from "react-redux";
import { Link, useNavigate } from "react-router-dom";


const CreateOrder = () => {
  const navigate = useNavigate();

  const [description, setDescription] = useState("");
  const [price, setPrice] = useState(0);
  const [startLocation, setStartLocation] = useState("");
  const [endLocation, setEndLocation] = useState("");
  const token = localStorage.getItem("token");

  const handleCreateOrder = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post(
        "http://localhost:8020/order/create",
        {
          description,
          start_location: startLocation,
          end_location: endLocation
        },
        {
          headers: {
            Authorization: `Bearer ${token}`
          }
        }
      );
      console.log("Order created successfully:");
      toast.success("Order created successfully");
      navigate("/");
    } catch (error) {
      console.error("Error creating order:", error);
      toast.error("Error creating order");
    }
  };

  return (
    <div className="flex flex-col justify-center sm:py-12">
      <div className="p-10 xs:p-0 mx-auto md:w-full md:max-w-md">
        <div className="bg-dark border border-gray-600 shadow w-full rounded-lg divide-y divide-gray-200">
          <form className="px-5 py-7" onSubmit={handleCreateOrder}>
            <div className="mb-4">
              <label className="font-semibold text-sm pb-1 block text-accent-content">
                Description
              </label>
              <input
                type="text"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                className="border rounded-lg px-3 py-2 mt-1 text-sm w-full"
              />
            </div>
      
            <div className="mb-4">
              <label className="font-semibold text-sm pb-1 block text-accent-content">
                Start Location
              </label>
              <input
                type="text"
                value={startLocation}
                onChange={(e) => setStartLocation(e.target.value)}
                className="border rounded-lg px-3 py-2 mt-1 text-sm w-full"
              />
            </div>
            <div className="mb-4">
              <label className="font-semibold text-sm pb-1 block text-accent-content">
                End Location
              </label>
              <input
                type="text"
                value={endLocation}
                onChange={(e) => setEndLocation(e.target.value)}
                className="border rounded-lg px-3 py-2 mt-1 text-sm w-full"
              />
            </div>
            <button
              type="submit"
              className="bg-blue-600 hover:bg-blue-500 text-white px-4 py-2 rounded-lg text-sm"
            >
              Create Order
            </button>
          </form>
        </div>
      </div>
    </div>
  );
};

export default CreateOrder;

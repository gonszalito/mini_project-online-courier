import React, { useEffect, useState } from "react";
import { SectionTitle, WishItem } from "../components";
import { useDispatch, useSelector } from "react-redux";
import axios from "axios"; // Import axios for making HTTP requests

const Wishlist = () => {
  const { wishItems } = useSelector((state) => state.wishlist);
  const dispatch = useDispatch();
  const [userData, setUserData] = useState(null); // State to store user data

  useEffect(() => {
    fetchUserData(); // Fetch user data when component mounts
  }, []); // Run effect only once when component mounts

  const getUsernameFromLocalStorage = () => {
    return localStorage.getItem('username');
  };

  const fetchUserData = async () => {
    const username = getUsernameFromLocalStorage();
    if (!username) {
      console.error('Username not found in localStorage');
      return;
    }

    const url = `http://localhost:8010/user/${username}`;

    try {
      const response = await axios.get(url);
      const userData = response.data.data; // Accessing the nested 'data' property
      console.log('User data:', userData);
      setUserData(userData); // Set user data to state
    } catch (error) {
      console.error('Error fetching user data:', error.message);
      // Handle error
    }
  };

  return (
    <>
      <SectionTitle title="" path="Profile" />
      {/* Display user data here */}
      <div className="max-w-7xl mx-auto">
        <div className="overflow-x-auto">
          {userData && (
            <div>
              <h2 className="text-xl font-semibold mb-2">User Profile</h2>
              <div>
                <p><strong>Name:</strong> {userData.name}</p>
                <p><strong>Email:</strong> {userData.email}</p>
                <p><strong>Username:</strong> {userData.username}</p>
                <p><strong>Role:</strong> {userData.role}</p>
              </div>
            </div>
          )}
          <table className="table">
            <thead>
              <tr>
                <th></th>
                <th className="text-accent-content">Name</th>
                <th className="text-accent-content">Size</th>
                <th className="text-accent-content">Action</th>
              </tr>
            </thead>
            <tbody>
              {wishItems.map((item, index) => (
                <WishItem item={item} key={index} counter={index} />
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </>
  );
};

export default Wishlist;

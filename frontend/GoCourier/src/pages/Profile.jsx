import React, { useEffect, useState } from "react";
import { SectionTitle } from "../components";
import axios from "axios";
import { useSelector } from "react-redux";
import { toast } from "react-toastify";
import { useNavigate } from "react-router-dom";

const Profile = () => {
  const [id, setId] = useState(localStorage.getItem("id"));
  const [userData, setUserData] = useState({});
  const loginState = useSelector((state) => state.auth.isLoggedIn);
  const wishItems = useSelector((state) => state.wishlist.wishItems);
  const [userFormData, setUserFormData] = useState({
    username: "",
    email: "",
    password: "",
  });
  const navigate = useNavigate();

  const getUsernameFromLocalStorage = () => {
    return localStorage.getItem('username');
  };

  const getUserData = async () => {
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

  useEffect(() => {
    if (loginState) {
      getUserData();
    } else {
      toast.error("You must be logged in to access this page");
      navigate("/");
    }
  }, []);

  const updateProfile = async (e) => {
    e.preventDefault();
    try{

      const getResponse = await axios(`http://localhost:8080/user/${id}`);
      const userObj = getResponse.data;

      // saljemo get(default) request
      const putResponse = await axios.put(`http://localhost:8080/user/${id}`, {
        id: id,
        name: userFormData.name,
        lastname: userFormData.lastname,
        email: userFormData.email,
        phone: userFormData.phone,
        adress: userFormData.adress,
        password: userFormData.password,
        userWishlist: await userObj.userWishlist
        //userWishlist treba da stoji ovde kako bi sacuvao stanje liste zelja
      });
      const putData = putResponse.data;
    }catch(error){
      console.log(error.response);
    }
  }

  return (
    <>
      <SectionTitle title="" path="User Profile" />
      <div className="max-w-7xl mx-auto flex justify-center items-center h-full">
        <div className="overflow-x-auto pb-10">
          {userData && (
            <div>
              {/* <h2 className="text-xl font-semibold mb-2">User Profile</h2> */}
              <div>
                <p><strong>Name:</strong> {userData.name}</p>
                <p><strong>Email:</strong> {userData.email}</p>
                <p><strong>Username:</strong> {userData.username}</p>
                <p><strong>Role:</strong> {userData.role}</p>
              </div>
            </div>
          )}
        </div>
      </div>

    </>
  );
};

export default Profile;

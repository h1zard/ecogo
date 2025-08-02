// src/pages/Items.jsx
import React, { useEffect, useState, useContext } from "react";
import { useNavigate } from "react-router-dom";
import {
  getItems,
  createCart,
  createOrder,
  getCarts,
  getOrders,
} from "../api";
import { toast } from "sonner";
import { AuthContext } from "../context/AuthContext";
import "../styles.css";

const Items = () => {
  const [items, setItems] = useState([]);
  const navigate = useNavigate();
  const { logout } = useContext(AuthContext); // ✅ using AuthContext

  useEffect(() => {
    fetchItems();
  }, []);

  const fetchItems = async () => {
    try {
      const res = await getItems();
      setItems(res.data);
    } catch (err) {
      toast.error("Failed to fetch items");
    }
  };

  const addToCart = async (itemId) => {
    try {
      await createCart({ item_id: itemId });
      toast.success("Item added to cart");
    } catch (err) {
      toast.error("Failed to add item to cart");
    }
  };

  const checkout = async () => {
    try {
      await createOrder({});
      toast.success("Order successful");
    } catch (err) {
      toast.error("Checkout failed");
    }
  };

  const showCart = async () => {
  try {
    const res = await getCarts();
    const items = res.data.Items.map(ci => `Item ID: ${ci.Item.ID}, Name: ${ci.Item.Name}`).join("\n");
    alert(items || "Your cart is empty");
  } catch (err) {
    toast.error("Failed to fetch cart");
  }
};


  const showOrders = async () => {
    try {
      const res = await getOrders();
      alert(res.data.map((o) => o.ID).join(", "));
    } catch (err) {
      toast.error("Failed to fetch orders");
    }
  };

  const handleLogout = () => {
    logout(); // ✅ from context
    toast.success("Logged out successfully");
    navigate("/login");
  };

  return (
    <div className="items-container">
      <div className="actions">
        <button onClick={checkout}>Checkout</button>
        <button onClick={showCart}>Cart</button>
        <button onClick={showOrders}>Order History</button>
        <button onClick={handleLogout}>Logout</button>
      </div>
      <h2>Items</h2>
      <div className="items-grid">
        {items.map((item) => (
          <div
            key={item.ID}
            className="item-card"
            onClick={() => addToCart(item.ID)}
          >
            <h4>{item.Name}</h4>
            <p>{item.Status}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Items;

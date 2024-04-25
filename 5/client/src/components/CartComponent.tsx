import React from 'react';
import { useCart } from '../contexts/ContextCart';
import { Product } from "../interfaces/Product";

const CartComponent: React.FC = () => {
  const { cart, cartTotalAmount, addProductToCart, removeProductFromCart } = useCart();

  return (
    <div>
      <h2>Shopping Cart</h2>
      {cart.length === 0 ? (
        <p>No products in the cart</p>
      ) : (
        <div>
          <table>
            <thead>
            <tr>
              <th>Name</th>
              <th>Description</th>
              <th>Price</th>
              <th>Quantity</th>
              <th>Add more</th>
              <th>Remove one</th>
            </tr>
            </thead>
            <tbody>
            {cart.map((product: Product) => (
              <tr key={product.id}>
                <td>{product.name}</td>
                <td>{product.description}</td>
                <td>{product.price}</td>
                <td>{product.quantity}</td>
                <td>
                  <button onClick={() => addProductToCart(product)}>+</button>
                </td>
                <td>
                  <button onClick={() => removeProductFromCart(product.id)}>-</button>
                </td>

              </tr>
            ))}
            </tbody>
          </table>
          <div>
            <p>Total Amount: {cartTotalAmount.toFixed(2)}</p>
          </div>
        </div>
      )}
    </div>
  );
};

export default CartComponent;

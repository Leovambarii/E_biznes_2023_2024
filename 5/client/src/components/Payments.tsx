import React, { useState } from 'react';
import { Payment } from '../interfaces/Payment';
import { useCart } from "../contexts/ContextCart";
import CartComponent from "./CartComponent";

const Payments: React.FC = () => {
  const initialPaymentState: Payment = {
    cardNumber: '',
    expiryDate: '',
    amount: 0,
  };
  const [payment, setPayment] = useState<Payment>(initialPaymentState);
  const [response, setResponse] = useState<string>('');
  const { cartTotalAmount, clearCart } = useCart();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    const parsedValue = name === 'amount' ? parseFloat(value) : value;
    setPayment(prevPayment => ({
      ...prevPayment,
      [name]: parsedValue,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      console.log(typeof payment.amount);
      const response = await fetch('http://localhost:8080/payment', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(payment),
      });
      if (response.ok) {
        const data = await response.json();
        clearCart();
        setResponse(JSON.stringify(data));
      } else {
        throw new Error('Failed to make payment.');
      }
    } catch (error) {
      console.error('Error making payment:', error);
      setResponse('Error: Unable to make payment. Please try again.');
    }
  };

  return (
    <div>
      <h2>Payment</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="cardNumber">Card Number:</label>
          <input
            type="text"
            id="cardNumber"
            name="cardNumber"
            value={payment.cardNumber}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor="expiryDate">Expiry Date:</label>
          <input
            type="text"
            id="expiryDate"
            name="expiryDate"
            value={payment.expiryDate}
            onChange={handleChange}
          />
        </div>
        <div>
          <p>Total Amount: {cartTotalAmount.toFixed(2)}</p>
        </div>
        <button type="submit">Make Payment</button>
      </form>
      {response && <div>Response: {response}</div>}

      <CartComponent />
    </div>
  );
};

export default Payments;

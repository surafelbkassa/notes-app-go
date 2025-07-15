import React from "react";

import React from 'react'

const signup = () => {
  return (
    <div>
        <form>
            <h1>Signup</h1>
            <label htmlFor="name">Name</label>
            <input
                placeholder="Enter Your name"
                type="text"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
             />
            <label htmlFor="email">Email</label>
            <input
                placeholder="Enter Your email"
                value={email}
                type="email"
                onChange={(e) => setEmail(e.target.value)}
                />
            <label htmlFor="password">Password</label>
            <input 
                placeholder="Enter Your password"
                value={password}
                type="password"
                onChange={(e) => setPassword(e.target.value)}
            />
        </form>          
    </div>
  )
}

export default signup

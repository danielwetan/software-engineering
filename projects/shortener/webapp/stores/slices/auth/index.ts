import { createSlice } from "@reduxjs/toolkit";

type AuthState = {
  data: Profile | null;
};

const initialState: AuthState = {
  data: null,
};

export const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    setData(state, action) {
      state.data = action.payload?.data;
    },
  },
});

export const { setData } = authSlice.actions;

export default authSlice.reducer;

import { createSlice } from "@reduxjs/toolkit";

type DialogDataState = {
  login: boolean;
  register: boolean;
};

type DialogState = {
  data: DialogDataState;
};

const initialState: DialogState = {
  data: {
    login: false,
    register: false,
  },
};

export const dialogSlice = createSlice({
  name: "dialog",
  initialState,
  reducers: {
    setDialog(state, action) {
      state.data = {
        ...state.data,
        ...action.payload,
      };
    },
  },
});

export const { setDialog } = dialogSlice.actions;

export default dialogSlice.reducer;

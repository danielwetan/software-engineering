import { combineReducers } from "redux";
import auth from "../slices/auth";
import dialog from "../slices/dialog";

const appReducer = combineReducers({
  authReducer: auth,
  dialogReducer: dialog,
});

export default appReducer;

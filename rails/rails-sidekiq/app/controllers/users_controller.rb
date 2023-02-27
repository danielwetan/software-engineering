class UsersController < ApplicationController
  def index
    response = {
      message: 'it works'
    }

    # execute background worker
    WelcomeEmailWorker.perform_async("Daniel")

    render json: response
  end
end

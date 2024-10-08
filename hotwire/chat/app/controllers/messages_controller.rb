class MessagesController < ApplicationController
  before_action :set_room, only: %i[ new create ]

  # GET /rooms/new
  def new
    @message = @room.messages.new
  end

  # POST /rooms or /rooms.json
  def create
    @message = @room.messages.create!(message_params)

    respond_to do |format|
      # comment this to prevent duplicate turbo stream after submit message using cable (websocket)
      # format.turbo_stream
      format.html { redirect_to @room }
    end
  end

  private
    def set_room
      @room = Room.find(params[:room_id])
    end

    # Only allow a list of trusted parameters through.
    def message_params
      params.require(:message).permit(:content)
    end
end

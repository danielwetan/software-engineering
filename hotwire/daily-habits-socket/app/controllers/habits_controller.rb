class HabitsController < ApplicationController
  before_action :set_habit 

  def show
  end

  def plus 
    @habit.update(count: @habit.count + 1)

    Turbo::StreamsChannel.broadcast_replace_to(
      "success_action_#{@habit.id}",
      target: "success-message",
      template: "habits/success_perform_plus",
      layout: false,
      locals: { habit: @habit }
    )

    render :result
  end

  def minus 
    @habit.update(count: @habit.count - 1)

    Turbo::StreamsChannel.broadcast_replace_to(
      "success_action_#{@habit.id}",
      target: "success-message",
      template: "habits/success_perform_minus",
      layout: false,
      locals: { habit: @habit }
    )
    render :result
  end

  private 
  def set_habit
    @habit = Habit.find_by(id: params['id'])
  end
end

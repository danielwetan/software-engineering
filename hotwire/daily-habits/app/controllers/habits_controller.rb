class HabitsController < ApplicationController
  before_action :set_habit 

  def show
    @habit = Habit.first
  end

  def plus 
    @habit.update(count: @habit.count + 1)
    redirect_to @habit
  end

  def minus 
    @habit.update(count: @habit.count - 1)
    redirect_to @habit
  end

  private 
  def set_habit
    @habit = Habit.find_by(id: params['id'])
  end
end

<%= turbo_frame_tag @habit do %>
  <div class="text-center font-bold text-gray-700 " id="habit-name">
    <%= @habit.name %>
  </div>

  <div class="mt-3 flex justify-center items-center space-x-5">
    <%= button_to '-', minus_habit_path(@habit), class: 'rounded-lg py-1 px-3 inline-block font-medium cursor-pointer shadow-lg bg-red-300' %>
    <div class="text-4xl font-bold"><%= @habit.count %></div>
    <%= button_to '+', plus_habit_path(@habit), class: 'rounded-lg py-1 px-3 inline-block font-medium cursor-pointer shadow-lg bg-green-300' %>
  </div>

  <div class="mt-3 p-2 flex justify-center space-x-1">
    <% @habit.count.times do %>
      <div class="inline-block border p-1 bg-green-400"></div>
    <% end %>
  </div>
<% end %>

class WelcomeEmailWorker
  include Sidekiq::Worker
  sidekiq_options retry: false

  def perform(name)
    p "Trying to sending email to #{name}"
    # simulate long processing job
    sleep(10.seconds)
    p "Sending email to #{name} success"
  end
end
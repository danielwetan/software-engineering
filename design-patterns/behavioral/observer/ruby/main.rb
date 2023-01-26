# The Subject interface declares a set of methods for managing subscribers.
class Subject
  # Attach an observer to the subject.
  def attach(observer)
    raise NotImplementedError, "#{self.class} has not implemented method '#{__method__}'"
  end

  # Detach an observer from the subject.
  def detach(observer)
    raise NotImplementedError, "#{self.class} has not implemented method '#{__method__}'"
  end

  # Notify all observers about an event.
  def notify
    raise NotImplementedError, "#{self.class} has not implemented method '#{__method__}'"
  end
end



# The Subject owns some important state and notifies observers when the state
# changes.
class ConcreteSubject < Subject
  # For the sake of simplicity, the Subject's state, essential to all
  # subscribers, is stored in this variable.
  attr_accessor :state

  # @!attribute observers
  # @return [Array<Observer>] attr_accessor :observers private :observers

  def initialize
    @observers = []
  end

  # List of subscribers. In real life, the list of subscribers can be stored
  # more comprehensively (categorized by event type, etc.).

  # @param [Obserser] observer
  def attach(observer)
    puts 'Subject: Attached an observer.'
    @observers << observer
  end

  # @param [Obserser] observer
  def detach(observer)
    @observers.delete(observer)
  end

  # The subscription management methods.

  # Trigger an update in each subscriber.
  def notify
    puts 'Subject: Notifying observers...'
    @observers.each { |observer| observer.update(self) }
  end
end





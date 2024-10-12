class Payment < ApplicationRecord
  # conflict with built in active record transaction
  # rename the relation to associated_transaction
  belongs_to :associated_transaction, class_name: "Transaction", foreign_key: "transaction_id"

  validates :method, presence: true
  validates :amount, numericality: { greater_than: 0 }
  validates :status, inclusion: { in: [ "pending", "paid", "failed" ] }
end

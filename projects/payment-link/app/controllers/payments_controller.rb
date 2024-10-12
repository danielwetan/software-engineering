class PaymentsController < ApplicationController
  def create
    @transaction = Transaction.find(params[:transaction_id])

    @payment = Payment.new(
      associated_transaction: @transaction,  # Use the renamed association
      method: "credit_card",
      amount: @transaction.amount,
      status: "paid"
    )

    @payment.save!
    @transaction.update!(status: "paid")
    redirect_to transactions_path, notice: "Payment was successfully processed."

  rescue StandardError => e
    puts "Got error"
    puts e
  end
end

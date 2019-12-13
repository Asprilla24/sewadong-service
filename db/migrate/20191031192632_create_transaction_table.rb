class CreateTransactionTable < ActiveRecord::Migration[5.2]
  def change
    create_table :transactions, id: false do |t|
      t.string :transaction_id, limit: 128, null: false, primary_key: true
      t.string :user_id, limit: 128, null: false, unique: true
      t.string :product_id, limit: 128, null: false, unique: true
      t.string :status, limit: 128, null: false
      t.datetime :date_from, null: false, :default => Time.now
      t.datetime :date_to, null: false, :default => Time.now
      t.decimal :total_price, null:false
      t.datetime :updated_at, null: false, :default => Time.now
      t.datetime :created_at, null: false, :default => Time.now
    end
  end
end

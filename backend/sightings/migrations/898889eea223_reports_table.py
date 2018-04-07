"""Reports table

Revision ID: 898889eea223
Revises: 
Create Date: 2018-04-07 20:56:57.661221

"""
from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision = '898889eea223'
down_revision = None
branch_labels = ('default',)
depends_on = None


def upgrade():
    # ### commands auto generated by Alembic - please adjust! ###
    op.create_table('reports',
    sa.Column('id', sa.BigInteger(), nullable=False),
    sa.Column('location', sa.Text(), nullable=True),
    sa.Column('shape', sa.Text(), nullable=True),
    sa.Column('duration', sa.Text(), nullable=True),
    sa.Column('description', sa.Text(), nullable=True),
    sa.Column('latitude', sa.Numeric(precision=2), nullable=True),
    sa.Column('longitude', sa.Numeric(precision=2), nullable=True),
    sa.Column('reported_at', sa.DateTime(), nullable=True),
    sa.Column('occurred_at', sa.DateTime(), nullable=True),
    sa.PrimaryKeyConstraint('id')
    )
    # ### end Alembic commands ###


def downgrade():
    # ### commands auto generated by Alembic - please adjust! ###
    op.drop_table('reports')
    # ### end Alembic commands ###